package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB
var bucketName = []byte("posts")

// Blog post model
type Post struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// Initialize BoltDB and create bucket
func initDB() {
	var err error

	db, err = bolt.Open("blog.db", 0600, nil)
	if err != nil {
		panic(fmt.Sprintf("Failed to open database: %v", err))
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to create bucket: %v", err))
	}
}

// Convert numeric IDs into zero-padded keys
func idToBytes(id uint64) []byte {
	return []byte(fmt.Sprintf("%05d", id))
}

// Extract ID from URL paths like:
// /posts/5
func getIDFromURL(path string) (uint64, error) {
	idStr := strings.TrimPrefix(path, "/posts/")
	return strconv.ParseUint(idStr, 10, 64)
}

// Helper for sending JSON responses
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}


// Create a new blog post
func createPost(w http.ResponseWriter, r *http.Request) {

	var post Post

	// TODO 1:
	// Decode the JSON request body into:
	// post
	//
	// If decoding fails:
	// return http.StatusBadRequest


	// TODO 2:
	// Validate:
	// - post.Title
	// - post.Content
	//
	// If either field is empty:
	// return http.StatusBadRequest


	// TODO 3:
	// Save the post into BoltDB
	//
	// Steps:
	//
	// - Start a db.Update() transaction
	// - Get the bucket
	//
	// - Generate a new ID
	//
	// - Assign:
	//   post.ID
	//   post.CreatedAt
	//
	// - Use:
	//   time.Now().Format("2006-01-02 15:04:05")
	//
	// - Marshal the post into JSON
	//
	// - Save this data


	// TODO 4:
	// Return:
	// http.StatusCreated
	//
	// Use:
	// writeJSON(...)
	//
	// Return the newly created post object.


}

// Fetch all blog posts
func listPosts(w http.ResponseWriter, r *http.Request) {

	var posts []Post

	// TODO 5:
	// Fetch all posts from BoltDB
	//
	// Steps:
	//
	// - Start a db.View() transaction
	// - Get the bucket
	// - Iterate the bucket
	//
	// Inside the loop:
	//
	// - Unmarshal each post
	// - Append it into posts slice
	//
	// IMPORTANT:
	// If db.View() returns an error:
	// return http.StatusInternalServerError


	// TODO 6:
	// If no posts exist:
	// initialize empty posts slice
	//
	// This ensures the API returns:
	// []
	// instead of:
	// null


	// TODO 7:
	// Return:
	// http.StatusOK
	//
	// Use:
	// writeJSON(...)
	//
	// Return the full posts slice.


}

// Fetch a single post by ID
func getPost(w http.ResponseWriter, r *http.Request) {

	id, err := getIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid Post ID", http.StatusBadRequest)
		return
	}

	var post Post
	var found bool

	// TODO 8:
	// Fetch the post from BoltDB
	//
	// Steps:
	//
	// - Start a db.View() transaction
	// - Fetch the post
	//
	// If data exists:
	//
	// - Unmarshal into:
	//   post
	//
	// - Set:
	//   found = true
	//
	// IMPORTANT:
	// If db.View() returns an error:
	// return http.StatusInternalServerError


	// TODO 9:
	// If the post does not exist:
	// return:
	// http.StatusNotFound


	// TODO 10:
	// Return:
	// http.StatusOK
	//
	// Use:
	// writeJSON(...)
	//
	// Return the fetched post.


}

// Update an existing post
func updatePost(w http.ResponseWriter, r *http.Request) {

	id, err := getIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid Post ID", http.StatusBadRequest)
		return
	}

	var updatedData Post

	// TODO 11:
	// Decode the request body into:
	// updatedData
	//
	// Validate:
	// - Title
	// - Content
	//
	// If validation fails:
	// return http.StatusBadRequest


	var existingPost Post
	var found bool

	// TODO 12:
	// Update the post in BoltDB
	//
	// Steps:
	//
	// - Start a db.Update() transaction
	// - Fetch the existing post
	//
	// If the post exists:
	//
	// - Set:
	//   found = true
	//
	// - Unmarshal into:
	//   existingPost
	//
	// - Update:
	//   existingPost.Title
	//   existingPost.Content
	//
	// IMPORTANT:
	// Preserve:
	// - existingPost.ID
	// - existingPost.CreatedAt
	//
	// - Marshal updated data
	// - Save this back into BoltDB


	if err != nil {
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	// TODO 13:
	// If the post was not found:
	// return http.StatusNotFound
	//
	// Otherwise:
	// return http.StatusOK
	//
	// Use:
	// writeJSON(...)
	//
	// Return the updated post.


}

// Delete a post by ID
func deletePost(w http.ResponseWriter, r *http.Request) {

	id, err := getIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid Post ID", http.StatusBadRequest)
		return
	}

	// TODO 14:
	// Delete the post from BoltDB
	//
	// Steps:
	//
	// - Start a db.Update() transaction
	// - Delete the post using the ID


	// TODO 15:
	// Return:
	// http.StatusNoContent

	
}


func main() {
	fmt.Println("=== Blogging Platform API ===")

	initDB()
	defer db.Close()

	mux := http.NewServeMux()

	// Serve frontend UI
	mux.Handle("/", http.FileServer(http.Dir(".")))

	// Collection routes
	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodPost:
			createPost(w, r)

		case http.MethodGet:
			listPosts(w, r)

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Single resource routes
	mux.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			getPost(w, r)

		case http.MethodPut:
			updatePost(w, r)

		case http.MethodDelete:
			deletePost(w, r)

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Listening on port 8080 ...")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}


