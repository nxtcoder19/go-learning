package main

//
//func main() {
//	// Create a new Fiber app
//	app := fiber.New()
//
//	// Define a route for the root endpoint "/"
//	app.Get("/", func(c *fiber.Ctx) error {
//		return c.SendString("Welcome to the Fiber HTTP Server!")
//	})
//
//	// Define a route with a parameter
//	app.Get("/hello/:name", func(c *fiber.Ctx) error {
//		name := c.Params("name") // Get the "name" parameter from the URL
//		return c.SendString("Hello, " + name + "!")
//	})
//
//	// Start the server on port 8080
//	app.Listen(":8080")
//}
