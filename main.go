package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Posts struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

//func main() {
//	http.HandleFunc("/fetch-post", fetchHandler)
//	log.Println("Starting server on port 8080")
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
//
//func fetchHandler(writer http.ResponseWriter, request *http.Request) {
//	log.Println("inside fetchHandler method")
//	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
//	if err != nil {
//		log.Println("error fetching posts: ", err)
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//	}
//	//log.Println("resp fetch posts: ", resp)
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			log.Println("error closing body: ", err)
//		}
//	}(resp.Body)
//
//	result, err := readData(resp)
//	if err != nil {
//		log.Println("error parsing body: ", err)
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//	}
//	userIdString := request.URL.Query().Get("userId")
//	if userIdString == "" {
//		http.Error(writer, "userId is required", http.StatusBadRequest)
//		return
//	}
//	userId, err := strconv.Atoi(userIdString)
//	if err != nil {
//		http.Error(writer, "userId must be an integer", http.StatusBadRequest)
//		return
//	}
//
//	log.Println("userId:", userId)
//
//	var userPost []Posts
//	for _, post := range result {
//		if post.UserId == userId {
//			userPost = append(userPost, post)
//		}
//	}
//	log.Println("userPost:", userPost)
//	response, err := json.Marshal(userPost)
//	if err != nil {
//		log.Println("error marshal: ", err)
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//	}
//	writer.Header().Set("Content-Type", "application/json")
//	writer.WriteHeader(http.StatusOK)
//	_, err = writer.Write(response)
//	if err != nil {
//		log.Println("error writing response: ", err)
//		return
//	}
//}
//
//func readData(resp *http.Response) ([]Posts, error) {
//	log.Println("reading data")
//	//body, err := io.ReadAll(resp.Body)
//	//if err != nil {
//	//	log.Println("error reading body: ", err)
//	//	return []Posts{}, err
//	//}
//	//err = json.Unmarshal(body, &result)
//	//if err != nil {
//	//	log.Println("error parsing body: ", err)
//	//	return result, err
//	//}
//	//log.Println("result: ", result)
//	var result []Posts
//	err := json.NewDecoder(resp.Body).Decode(&result)
//	if err != nil {
//		return nil, err
//	}
//	return result, nil
//}

func main() {
	http.HandleFunc("/fetch-post", handleFetchUser)
	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error: unable to start sever due to: \n", err)
		return
	}
}

func handleFetchUser(writer http.ResponseWriter, request *http.Request) {
	city := request.URL.Query().Get("city")
	userDataResp, err := getDataFromEndpoint("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Println("error: unable to fetch user data due to: \n", err)
		http.Error(writer, "unable to fetch user data", http.StatusInternalServerError)
		return
	}
	userData, err := readUsersData(userDataResp.Body)
	if err != nil {
		log.Println("error: unable to read user data due to: \n", err)
		http.Error(writer, "unable to read user data", http.StatusInternalServerError)
		return
	}
	userId := searchInUserData(userData, city)
	if userId < 1 {
		log.Printf("error: unable to find user with given city %s: ", city)
		http.Error(writer, "unable to find user with given city", http.StatusNotFound)
		return
	}

	postsFromEndpoint, err := getDataFromEndpoint("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Println("error: unable to fetch post data from endpoint due to: \n", err)
		http.Error(writer, " error while fetching posts for user", http.StatusInternalServerError)
		return
	}

	userPosts, err := readUsersPosts(postsFromEndpoint.Body)
	if err != nil {
		log.Println("error: unable to read user posts due to: \n", err)
		http.Error(writer, "unable to read user posts", http.StatusInternalServerError)
		return
	}
	defer func() {
		err = postsFromEndpoint.Body.Close()
		if err != nil {
			log.Println("error: unable to close response body due to: \n", err)
		}
		err = userDataResp.Body.Close()
		if err != nil {
			log.Println("error: unable to close response body due to: \n", err)
		}
	}()

	posts := fetchUserPostById(userId, userPosts)
	if len(posts) == 0 {
		log.Println("no posts found for the user from: ", city)
	}
	resp, err := json.Marshal(&posts)
	if err != nil {
		log.Println("error unable to marshal response due to :\n", err)
		http.Error(writer, "error unable to marshal response", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	//writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(resp)
	if err != nil {
		log.Println("error writing response due to: ", err)
		return
	}
}

func fetchUserPostById(userId int, posts []Posts) []Posts {
	var userPosts []Posts
	for _, post := range posts {
		if post.UserId == userId {
			userPosts = append(userPosts, post)
		}
	}
	return userPosts
}

func searchInUserData(data []User, city string) int {
	for _, user := range data {
		if user.Address.City == city {
			return user.Id
		}
	}
	return -1
}

func readUsersData(respBody io.Reader) ([]User, error) {
	var userData []User
	err := json.NewDecoder(respBody).Decode(&userData)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func readUsersPosts(respBody io.Reader) ([]Posts, error) {
	var userPosts []Posts
	err := json.NewDecoder(respBody).Decode(&userPosts)
	if err != nil {
		return nil, err
	}
	return userPosts, nil
}

func getDataFromEndpoint(url string) (*http.Response, error) {
	respFromEndpoint, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return respFromEndpoint, nil
}
