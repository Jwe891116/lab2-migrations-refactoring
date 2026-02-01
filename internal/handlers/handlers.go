package handlers

import (
	"net/http"
)

// Home handles requests to the root path
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Lab #1: Building my first web application.\nMy name is Joshua Emmanuel.\nI will be doing the Medical Appointment Scheduling System for my Semester Project.\nReason, I believe a similar system could be implemented for the Medical Department for the Belize Coast Guard, my current place of employment.\n"))
}

// About handles requests to the /about path
func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, I'm Joshua Emmanuel, I'm a member of the Belize Coast Guard with seventeen years of experience in Operations and minor IT work.\n"))
}

// Contact handles requests to the /contact path
func Contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cell #: +501 000-0000\nEmail: 2005112848@ub.edu.bz\nGitHub: github.com/JWEmmanuel89\n"))
}

// Calculate handles requests to the /calculate path
func Calculate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("5+5*5-5 = 25\n"))
}
