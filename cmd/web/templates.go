package main

import "github.com/vysmv/snippetbox-public/internal/models"

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the project progresses.
type templateData struct {
    Snippet models.Snippet
}