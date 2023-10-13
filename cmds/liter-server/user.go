
package main

import (
	"bufio"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"io"
	"os"
	"sort"
	"sync"
)

var (
	ErrUserExists = errors.New("Username is already exists")
)

type User struct {
	Name     string `json:"name"`
	// password is encoded as sha256
	Password string `json:"password"`
}

func (u *User)SetPassword(pswd string){
	u.Password = asSha256(pswd)
}

func (u *User)CheckPassword(pswd string)(ok bool){
	return subtle.ConstantTimeCompare(([]byte)(u.Password), ([]byte)(asSha256(pswd))) == 1
}

type UserStorage struct {
	file  string
	mux   sync.RWMutex
	users []*User
}

func NewUserStorage(file string)(s *UserStorage){
	s = &UserStorage{
		file: file,
	}
	return
}

func (s *UserStorage)Len()(int){
	return len(s.users)
}

func (s *UserStorage)Load()(err error){
	s.mux.Lock()
	defer s.mux.Unlock()

	fd, err := os.Open(s.file)
	if err != nil {
		return
	}
	defer fd.Close()
	r := bufio.NewReader(fd)
	users := make([]*User, 0, 3)
	decoder := json.NewDecoder(r)
	for {
		u := new(User)
		if err = decoder.Decode(u); err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
				break
			}
			return
		}
		users = append(users, u)
	}
	sort.Slice(users, func(i, j int)(bool){
		return users[i].Name < users[j].Name
	})
	s.users = users
	return
}

func (s *UserStorage)Save()(err error){
	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.save()
}

func (s *UserStorage)save()(err error){
	fd, err := os.Create(s.file)
	if err != nil {
		return
	}
	defer fd.Close()
	w := bufio.NewWriter(fd)
	encoder := json.NewEncoder(fd)
	for _, u := range s.users {
		if err = encoder.Encode(u); err != nil {
			return
		}
	}
	if err = w.Flush(); err != nil {
		return
	}
	return
}

func (s *UserStorage)searchUser(name string)(u *User, i int){
	i = sort.Search(len(s.users), func(i int)(bool){
		return s.users[i].Name >= name
	})
	if i >= len(s.users) {
		return
	}
	if u = s.users[i]; u.Name != name {
		u = nil
		return
	}
	return
}

func (s *UserStorage)GetUser(name string)(u *User){
	s.mux.RLock()
	defer s.mux.RUnlock()

	u, _ = s.searchUser(name)
	return
}

func (s *UserStorage)AddUser(u *User)(err error){
	s.mux.Lock()
	defer s.mux.Unlock()

	o, i := s.searchUser(u.Name)
	if o != nil {
		return ErrUserExists
	}

	old := s.users
	s.users = insert(s.users, i, u)
	if err = s.save(); err != nil {
		s.users = old
		return
	}
	return
}

func (s *UserStorage)DelUser(name string)(err error){
	s.mux.Lock()
	defer s.mux.Unlock()

	if u, i := s.searchUser(name); u != nil {
		remove(s.users, i)
		return s.save()
	}
	return
}
