<a href="https://resynced.design" align="center">
    <img src="https://r2.interrupted.me/uploads/5IWuDvqm.png" align="center" />
</a>

<h1 align="center">Resynced Uploader Template</h1>

<p align="center">Star for a cookie 🍪</p>

---

Welcome to your **Resynced Uploader Project**! This guide will walk you through setting up the project on your local computer and deploying it to the web. We’ve made this as easy as possible, so you don’t need any coding knowledge—just follow the steps!

---

## 🌟 Project Highlights

Your project includes:

- **🚀 Performance** – Great performance with Gofiber
- **🔗 Cloudflare R2** – Ready for handling file uploads

*Created with ❤️ by [Resynced Design](https://resynced.design/)!*

---

## 🚀 Getting Started

### 1. Setting Up Locally

#### Requirements

- **Go** (1.23 or newer)
- **Git** (if you haven’t installed it, you can [download Git here](https://git-scm.com/))

#### Steps

1. **Download the Project Files**
   - Download the files from the link we provided or from GitHub.
   - Save them in a new folder on your computer.

2. **Open a Terminal**
   - Go to the folder where your project files are saved. 
   - **On Windows**: Right-click in the folder and select "Open in Terminal."
   - **On Mac**: Open Terminal, type `cd`, then drag the folder into the Terminal window to set your location.

3. **Edit the env**
   - Copy the `.env.example` file to `.env` and fill in the required fields.
   - To get your Cloudflare R2 credentials, follow this [tutorial](https://developers.cloudflare.com/r2/api/s3/tokens/).

4. **Start the Project**
   - In the Terminal, type:
     ```bash
     go run src/main.go
     ```
   - The webserver will start, and you can test it with Insomnia or Postman. Or by visiting `http://localhost:<port-you-specified>`.

---

### 2. Hosting the Project

If you have your own VPS (like DigitalOcean, AWS, or any other hosting service), you can deploy your project there. Follow these steps:

#### Prerequisites

- A VPS with root or sudo access.
- Go installed on the server. You can install it using:
    ```bash
    sudo apt update
    sudo apt install golang
    ```
- An HTTP server like Nginx or Apache for reverse proxy and domain handling.

#### Steps

1. **Upload Your Files**
   - Use SCP, rsync, or FTP software (e.g., FileZilla) to transfer your project files to the server.
   - Example using scp:
        ```bash
        scp -r /path/to/your/project root@your-server-ip:/path/to/your/project
        ```

2. **Set Up Environment Variables**
   - Navigate to your project directory on the server.
   - Copy the `.env.example` file to `.env` and configure it with your Cloudflare R2 credentials and other settings.

3. **Start the Application**
    - Run the following command to start the application:
        ```bash
        go run src/main.go
        ```
    - For production, use a process manager like PM2 or systemd to keep the app running:
        - ***Using PM2***:
            ```bash
            sudo npm install -g pm2
            pm2 start src/main.go --name resynced-uploader
            pm2 save
            pm2 startup
            ```
        - ***Using systemd***: Create a service file at `/etc/systemd/system/resynced-uploader.service`:
            ```ini
            [Unit]
            Description=Resynced Uploader Service
            After=network.target

            [Service]
            User=your-username
            WorkingDirectory=/path/to/project
            ExecStart=/usr/local/go/bin/go run src/main.go
            Restart=always

            [Install]
            WantedBy=multi-user.target
            ```
            Start and enable the service:
            ```bash
            sudo systemctl daemon-reload
            sudo systemctl start resynced-uploader
            sudo systemctl enable resynced-uploader
            ```

4. **Set Up a Reverse Proxy**
    - Install and configure **Nginx** or **Apache** to route traffic to your GoFiber app.
    - Example for **Nginx**:
        ```nginx
        server {
            listen 80;
            server_name yourdomain.com;

            location / {
                proxy_pass http://127.0.0.1:<your-app-port>;
                proxy_http_version 1.1;
                proxy_set_header Upgrade $http_upgrade;
                proxy_set_header Connection 'upgrade';
                proxy_set_header Host $host;
                proxy_cache_bypass $http_upgrade;
            }
        }
        ```
    - Restart Nginx:
        ```bash
        sudo systemctl restart nginx
        ```

5. **Enable HTTPS (Optional but Recommended)**
    - Use **Certbot** to install an SSL certificate for your domain:
        ```bash
        sudo apt install certbot python3-certbot-nginx
        sudo certbot --nginx -d yourdomain.com
        ```

6. **Done!**
    - Visit your domain in a browser to see your Resynced Uploader project live!

## ⚙️ Need Help?

If you need any changes or run into any issues, reach out to us at [Resynced Design](https://resynced.design/)! We’re here to help you make the most of your new site. Good luck, and enjoy! 🎉
