package MusicEntry

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMusicManager_Add(t *testing.T) {
	type fields struct {
		musics []MusicEntry
	}
	type args struct {
		music *MusicEntry
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"first", fields{[]MusicEntry{{"first", "first", "first", "first", "MP3"}}}, args{&MusicEntry{
			Id:     "first",
			Name:   "first",
			Artist: "first",
			Source: "first",
			Type:   "MP3",
		}}},
	}
	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMusicManager()
			m.Add(tt.args.music)
			ast := assert.New(t)
			ast.Equal(m.musics[index].Id, tt.args.music.Id)
			ast.Equal(m.musics[index].Type, tt.args.music.Type)
		})
	}
}

//func TestMusicManager_Find(t *testing.T) {
//	type fields struct {
//		musics []MusicEntry
//	}
//	type args struct {
//		name string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *MusicEntry
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &MusicManager{
//				musics: tt.fields.musics,
//			}
//			if got := m.Find(tt.args.name); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Find() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestMusicManager_Get(t *testing.T) {
//	type fields struct {
//		musics []MusicEntry
//	}
//	type args struct {
//		Index int
//	}
//	tests := []struct {
//		name      string
//		fields    fields
//		args      args
//		wantMusic *MusicEntry
//		wantErr   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &MusicManager{
//				musics: tt.fields.musics,
//			}
//			gotMusic, err := m.Get(tt.args.Index)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotMusic, tt.wantMusic) {
//				t.Errorf("Get() gotMusic = %v, want %v", gotMusic, tt.wantMusic)
//			}
//		})
//	}
//}
//
//func TestMusicManager_Len(t *testing.T) {
//	type fields struct {
//		musics []MusicEntry
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &MusicManager{
//				musics: tt.fields.musics,
//			}
//			if got := m.Len(); got != tt.want {
//				t.Errorf("Len() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestMusicManager_Remove(t *testing.T) {
//	type fields struct {
//		musics []MusicEntry
//	}
//	type args struct {
//		Index int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *MusicEntry
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &MusicManager{
//				musics: tt.fields.musics,
//			}
//			if got := m.Remove(tt.args.Index); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Remove() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
func TestNewMusicManager(t *testing.T) {
	var tests []struct {
		name string
		want *MusicManager
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMusicManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMusicManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
