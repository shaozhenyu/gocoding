package main

import (
	"fmt"
	"sort"
)

func main() {
	t := Constructor()
	t.PostTweet(1, 5)
	fmt.Println(t.GetNewsFeed(1))
	t.Follow(1, 2)
	t.PostTweet(2, 6)
	fmt.Println(t.GetNewsFeed(1))
	t.Unfollow(1, 2)
	fmt.Println(t.GetNewsFeed(1))
	fmt.Println(t.users[1])
}

type Twitter struct {
	increaseID int
	users map[int]*User
}

type User struct {
	uid       int
	followees map[int]struct{}
	articles  []article
}

func newUser(userId int) *User {
	return  &User{
		uid: userId,
		followees: make(map[int]struct{}),
		articles: make([]article, 0, 16),
	}
}

type article struct {
	id   int
	time int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		users: make(map[int]*User),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.users[userId]; ok {
		articles := this.users[userId].articles
		articles = append(articles, article{tweetId, this.increaseID})
		if len(articles) > 10 {
			articles = articles[1:]
		}
		this.users[userId].articles = articles
	} else {
		user := newUser(userId)
		user.articles = append(user.articles, article{tweetId, this.increaseID})
		this.users[userId] = user
	}
	this.increaseID++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	ret := make([]int, 0)
	articles := make([]article, 0)
	user, ok := this.users[userId]
	if !ok {
		return ret
	}
	articles = append(articles, user.articles...)
	fmt.Println("xx: ", articles)
	for followeeId := range user.followees {
		if followeer, ok := this.users[followeeId]; ok {
			articles = append(articles, followeer.articles...)
		}
	}
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].time > articles[j].time
	})
	for i := 0; i < len(articles) && i < 10; i++ {
		ret = append(ret, articles[i].id)
	}
	return ret
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if _, ok := this.users[followerId]; !ok {
		this.users[followerId] = newUser(followerId)
	}
	if _, ok := this.users[followeeId]; !ok {
		this.users[followeeId] = newUser(followeeId)
	}
	this.users[followerId].followees[followeeId] = struct{}{}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if _, ok := this.users[followerId]; !ok {
		this.users[followerId] = newUser(followerId)
	}
	if _, ok := this.users[followeeId]; !ok {
		this.users[followeeId] = newUser(followeeId)
	}
	if _, ok := this.users[followerId].followees[followeeId]; ok {
		delete(this.users[followerId].followees, followeeId)
	}
}
