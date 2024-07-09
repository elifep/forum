package handlers

import (
	"net/http"
)

func SetupRoutes() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view_post", viewPostHandler)
	http.HandleFunc("/submit_comment", submitCommentHandler)
	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/vote", voteHandler)
	http.HandleFunc("/vote_comment", voteCommentHandler)
	http.HandleFunc("/upload_movies", uploadMoviesHandler)
	http.HandleFunc("/upload_turkish", uploadTurkishHandler)
	http.HandleFunc("/upload_science", uploadScienceHandler)
	http.HandleFunc("/upload_food", uploadFoodHandler)
	http.HandleFunc("/upload_technology", uploadTechnologyHandler)
	http.HandleFunc("/upload_health", uploadHealthHandler)
	http.HandleFunc("/movies", moviesHandler)
	http.HandleFunc("/turkish", turkishHandler)
	http.HandleFunc("/science", scienceHandler)
	http.HandleFunc("/food", foodHandler)
	http.HandleFunc("/technology", technologyHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/change_password", ChangePasswordHandler)
	http.HandleFunc("/password_change", ServePasswordChangePage)
	http.HandleFunc("/userprofile", UserProfileHandler)
	http.HandleFunc("/auth/google/login", handleGoogleLogin)
	http.HandleFunc("/auth/google/callback", handleGoogleCallback)
	http.HandleFunc("/auth/github", handleGitHubLogin)
	http.HandleFunc("/auth/github/callback", handleGitHubCallback)
	http.HandleFunc("/auth/facebook", handleFacebookLogin)
	http.HandleFunc("/auth/facebook/callback", handleFacebookCallback)
	http.HandleFunc("/update_view_count", viewCountHandler)
	http.HandleFunc("/deletePost", deletePostHandler)
	http.HandleFunc("/deleteComment", deleteCommentHandler)
}
