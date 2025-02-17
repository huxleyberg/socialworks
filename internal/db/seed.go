package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/huxleyberg/socialworks/internal/models"
	"github.com/huxleyberg/socialworks/internal/posts"
	"github.com/huxleyberg/socialworks/internal/users"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var usernames = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi",
	"ivan", "judy", "karl", "laura", "mallory", "nina", "oscar", "peggy",
	"quinn", "rachel", "steve", "trent", "ursula", "victor", "wendy", "xander",
	"yvonne", "zack", "amber", "brian", "carol", "doug", "eric", "fiona",
	"george", "hannah", "ian", "jessica", "kevin", "lisa", "mike", "natalie",
	"oliver", "peter", "queen", "ron", "susan", "tim", "uma", "vicky",
	"walter", "xenia", "yasmin", "zoe",
}

var titles = []string{
	"The Power of Habit", "Embracing Minimalism", "Healthy Eating Tips",
	"Travel on a Budget", "Mindfulness Meditation", "Boost Your Productivity",
	"Home Office Setup", "Digital Detox", "Gardening Basics",
	"DIY Home Projects", "Yoga for Beginners", "Sustainable Living",
	"Mastering Time Management", "Exploring Nature", "Simple Cooking Recipes",
	"Fitness at Home", "Personal Finance Tips", "Creative Writing",
	"Mental Health Awareness", "Learning New Skills",
}

var contents = []string{
	"In this post, we'll explore how to develop good habits that stick and transform your life.",
	"Discover the benefits of a minimalist lifestyle and how to declutter your home and mind.",
	"Learn practical tips for eating healthy on a budget without sacrificing flavor.",
	"Traveling doesn't have to be expensive. Here are some tips for seeing the world on a budget.",
	"Mindfulness meditation can reduce stress and improve your mental well-being. Here's how to get started.",
	"Increase your productivity with these simple and effective strategies.",
	"Set up the perfect home office to boost your work-from-home efficiency and comfort.",
	"A digital detox can help you reconnect with the real world and improve your mental health.",
	"Start your gardening journey with these basic tips for beginners.",
	"Transform your home with these fun and easy DIY projects.",
	"Yoga is a great way to stay fit and flexible. Here are some beginner-friendly poses to try.",
	"Sustainable living is good for you and the planet. Learn how to make eco-friendly choices.",
	"Master time management with these tips and get more done in less time.",
	"Nature has so much to offer. Discover the benefits of spending time outdoors.",
	"Whip up delicious meals with these simple and quick cooking recipes.",
	"Stay fit without leaving home with these effective at-home workout routines.",
	"Take control of your finances with these practical personal finance tips.",
	"Unleash your creativity with these inspiring writing prompts and exercises.",
	"Mental Health is just as important as physical health. Learn how to take care of your mind.",
	"Learning new skills can be fun and rewarding. Here are some ideas to get you started.",
}

var tags = []string{
	"Self Improvement", "Minimalism", "Health", "Travel", "Mindfulness",
	"Productivity", "Home Office", "Digital Detox", "Gardening", "DIY",
	"Yoga", "Sustainability", "Time Management", "Nature", "Cooking",
	"Fitness", "Personal Finance", "Writing", "Mental Health", "Learning",
}

var comments = []string{
	"Great post! Thanks for sharing.",
	"I completely agree with your thoughts.",
	"Thanks for the tips, very helpful.",
	"Interesting perspective, I hadn't considered that.",
	"Thanks for sharing your experience.",
	"Well written, I enjoyed reading this.",
	"This is very insightful, thanks for posting.",
	"Great advice, I'll definitely try that.",
	"I love this, very inspirational.",
	"Thanks for the information, very useful.",
}

// Seed populates the database with initial data
func Seed(commentsRepo posts.CommentRepository, postsRepo posts.PostRepository,
	usersRepo users.UserRepository) {
	ctx := context.Background()

	// Generate and insert users
	users := generateUsers(100)
	for _, user := range users {
		if err := usersRepo.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	// Generate and insert posts
	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := postsRepo.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	// Generate and insert comments
	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := commentsRepo.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

// generateUsers creates a list of unique users
func generateUsers(num int) []*models.User {
	users := make([]*models.User, 0, num)
	usedEmails := make(map[string]bool)
	usedUsernames := make(map[string]bool)

	for i := 0; len(users) < num; i++ {
		username := usernames[i%len(usernames)] + fmt.Sprintf("%d", i)
		email := username + fmt.Sprintf("%d", i) + "@example.com"

		// Ensure uniqueness
		if usedEmails[email] || usedUsernames[username] {
			continue
		}

		usedEmails[email] = true
		usedUsernames[username] = true

		user := &models.User{
			Username:  username,
			Email:     email,
			Password:  "123123",
			CreatedAt: time.Now(),
		}
		users = append(users, user)
	}

	return users
}

// generatePosts creates a list of posts associated with users
func generatePosts(num int, users []*models.User) []*models.Post {
	posts := make([]*models.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &models.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

// generateComments creates a list of comments associated with posts and users
func generateComments(num int, users []*models.User, posts []*models.Post) []*models.Comment {
	cms := make([]*models.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &models.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
