🎨 Thesis: Optimizing Image Performance on SaaS Landing Pages
This repository explores the effects of image loading strategies on landing page performance, providing a comprehensive environment to test and analyze image optimization in a SaaS context.

🌐 Features
Smart Image Throttling: Simulates load under real-world conditions to assess performance impact.
Modular Go & JavaScript Integration: Flexible back-end and front-end testing framework.
In-Depth Performance Analytics: Provides data on load times, user engagement, and conversion.
🚀 Getting Started
1. Clone the Repository
Clone the repository locally to access all files and code needed to run the tests.

bash
Copy code
git clone https://github.com/sokennguyen/thesis
cd thesis
2. Install Dependencies
This project relies on Go modules. Run the following command to install dependencies:

bash
Copy code
go mod download
3. Run the Server
Start the Go server locally to make the app accessible on a specified port:

bash
Copy code
go run main.go
By default, this will serve the app on localhost:PORT. Replace PORT with your preferred port number.

🖥️ Configuring Nginx with Throttling
To simulate various network conditions and better observe loading behavior, you can set up Nginx as a reverse proxy with throttling capabilities. This enables you to limit the bandwidth for testing, simulating slower network conditions that users might experience.

1. Install Nginx
First, ensure that Nginx is installed on your server. Run the following commands to update your package list and install Nginx:

bash
Copy code
sudo apt update
sudo apt install nginx
2. Configure Nginx as a Reverse Proxy with Throttling
Open the default Nginx configuration file:
bash
Copy code
sudo nano /etc/nginx/sites-available/default
Inside the server block, add a location block for your app. The following configuration sets Nginx to proxy requests to your Go application while limiting the download rate:
nginx
Copy code
server {
    listen 80;
    server_name your_domain.com;  # Update with your domain or IP

    location / {
        proxy_pass http://localhost:PORT;  # Ensure PORT matches your Go server's port
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # Throttle download rate to 200 KB/s for testing
        proxy_limit_rate 200k;
    }
}
Here, the proxy_pass line tells Nginx to forward requests to your Go server, while proxy_limit_rate restricts the speed at which data is sent to clients, simulating slower connections.
3. Test and Restart Nginx
After saving your configuration, test the Nginx configuration for syntax errors:

bash
Copy code
sudo nginx -t
If there are no errors, restart Nginx to apply changes:

bash
Copy code
sudo systemctl restart nginx
Your Nginx server is now set up to forward requests to your application and throttle image loading speeds. This is essential for observing how different loading speeds affect user interaction and conversion rates on SaaS landing pages.

📊 Usage and Analysis
Once the Nginx proxy is configured and your app is running, access your site through your server’s IP or domain to observe real-time throttling effects. This setup will allow you to monitor image loading performance, assess user experience, and gather data for your thesis.

📜 License
This project is licensed under the MIT License, making it open for community contributions.


