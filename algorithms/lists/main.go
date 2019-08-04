package main

import "fmt"

type Feed struct {
	length int
	start  *Post
	end    *Post
}
type Post struct {
	body        string
	publishDate int64
	next        *Post
}

func (f *Feed) Append(newPost *Post) {
	if f.length == 0 {
		// its a first post
		f.start = newPost
		f.end = newPost
	} else {
		//append
		lastPost := f.end
		lastPost.next = newPost
		f.end = newPost
	}
	f.length++
}

func (f *Feed) Display(post *Post) {
	for post != nil {
		fmt.Printf("%v %v\n", post.body, post.publishDate)
		post = post.next
	}

}

// TODO fix remove update the main() with change if its first node
func (f *Feed) Remove(publishDate int64) {
	if f.length == 0 {
		fmt.Println("Error:Remove fails the list is empty")
		return
	}
	var prevPost *Post
	currentPost := f.start

	if f.start.publishDate == publishDate {
		prevPost = currentPost.next
		f.length--
		return
	}

	for currentPost.publishDate != publishDate {
		if currentPost.next == nil {
			fmt.Println("Error:Remove failed, no such post found")
		}

		prevPost = currentPost
		currentPost = currentPost.next
	}
	prevPost.next = currentPost.next
	f.length--
}
func main() {
	f := &Feed{}
	p := &Post{}

	f.Remove(1)

	for i := 0; i < 5; i++ {
		p1 := Post{body: "My  post", publishDate: int64(i)}
		f.Append(&p1)
	}
	p = f.start
	fmt.Println("Display the list of feeds..")
	f.Display(p)
	fmt.Printf("Length: %v\n", f.length)

	f.Remove(f.start.publishDate)
	f.Display(p)
}
