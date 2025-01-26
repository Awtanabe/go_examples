package map_test

import "testing"

func TestSearch(t *testing.T) {
	assrtStrings := func(t *testing.T, got, want string) {
		 t.Helper()
		if got != want {
			t.Errorf(" got %q want %q", got, want)
		}
	
	}


	t.Run("known key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assrtStrings(t, got, want)	
	})

	t.Run("unknown key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("aaa")
		want := "not found"

		if err == nil {
			t.Error("want error")
		}

		assrtStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {

	t.Run("追加成功", func(t *testing.T){
		want := "val"
		d := Dictionary{}
		d.Add("key", want)
		val, _ := d.Search("key")
		if val != want {
			t.Errorf("got %q want %q", val, want)
		}	
	})

	t.Run("追加失敗", func(t *testing.T){
		d := Dictionary{}
		err := d.Add("", "val")
		if err.Error() != "不足" {
			t.Errorf("got %q want %q", err.Error(), "不足")
		}	
	})
	// assrtStrings(t, err.Error(), want)
}
