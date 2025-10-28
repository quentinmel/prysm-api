package database

import (
    "log"
    "os"

    "github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

func InitSupabase() {
    url := os.Getenv("SUPABASE_URL")
    key := os.Getenv("SUPABASE_KEY")

    if url == "" || key == "" {
        log.Fatal("SUPABASE_URL and SUPABASE_KEY must be set")
    }

    var err error
    Client, err = supabase.NewClient(url, key, nil)
    if err != nil {
        log.Fatalf("Failed to initialize Supabase client: %v", err)
    }

    log.Println("✅ Supabase client initialized")
}