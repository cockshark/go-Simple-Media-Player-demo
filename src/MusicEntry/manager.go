package MusicEntry

import (
	"errors"
	"fmt"
)

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

// NewMusicManager 实现构造/初始化函数
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

/* 实现音乐库的管理功能 增删改 。。。*/

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(Index int) (music *MusicEntry, err error) {
	if Index < 0 || Index >= len(m.musics) { // make stronger please
		return nil, errors.New("index out of range") // 错误不应该首字母大写
	}
	return &m.musics[Index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, music := range m.musics {
		if music.Name == name {
			return &music
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(Index int) *MusicEntry {
	if Index < 0 || Index >= len(m.musics) {
		return nil
	}

	removeMusic := &m.musics[Index]
	m.musics = append(m.musics[:Index], m.musics[Index+1:]...)
	return removeMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if name == "" {
		fmt.Println("the name can not be '' ")
		return nil
	}

	for i, music := range m.musics {
		if music.Name == name {
			//removeMusic := &music
			//m.musics = append(m.musics[:i], m.musics[i+1:]...)
			//return removeMusic
			return m.Remove(i)
		}
	}
	return nil
}
