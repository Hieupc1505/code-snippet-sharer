package sup

import (
	"github.com/nedpals/supabase-go"
)

// Create a client for the database
func CreateClient() *supabase.Client {
	supabaseUrl := "https://wlgtdhyfabhamdoyprik.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6IndsZ3RkaHlmYWJoYW1kb3lwcmlrIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MzkzNzc1NTQsImV4cCI6MjA1NDk1MzU1NH0.qjDCB_qm9cS9O2xjCBzNA4m4FnhRT-F5OQie3IRRChY"
	return supabase.CreateClient(supabaseUrl, supabaseKey)
}
