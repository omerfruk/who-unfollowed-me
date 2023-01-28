# who-unfollowed-you

This app will help you to find out who unfollowed you on GitHub.

## How to use

1. Clone this repository:
    
    Run `git clone https://github.com/omerfruk/who-unfollowed-you.git`

2. Install dependencies:
    
    Run `go mod download`

3. Run the app:
    
    Run `go run main.go`

4. Enter your username:
    
    Enter your GitHub username.

5. (optional) Do you want to save the result to a file?
    
    Enter `y` or `n`.
    The app will print the result to the console and if you choose to save the result to a file, it will create a file named `result.txt` in the same directory.

6. End of the app:
    
    The app will print the unfollowers to the save the result to a file, it will create a file named `unfollowers.html` in the same directory.

NOTE 
1. If you want to use the app with a different username, You should know that `result.txt` and `unfollowers.html` files will be removed.
2. If You use a lots of GitHub API requests, you may get a `403 Forbidden` error. I will fix it if i want.
    

