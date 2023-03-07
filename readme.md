If you are looking to deploy chatgpt-demo on your server and don't know how to do it, this project can help you out!

Implementation Principle

![https://user-images.githubusercontent.com/34591322/223335950-8af037e8-d3aa-4ea5-ac3d-b5c281dcdc96.png](https://user-images.githubusercontent.com/34591322/223335950-8af037e8-d3aa-4ea5-ac3d-b5c281dcdc96.png)

Installation
If you have a Go environment, you can install it directly using:

    go get github.com/clearcodecn/chatgpt-demo-proxy
    
Alternatively, you can download binaries from the release page.

How to Use

1. [Deploy chatgpt-demo by Vercel](https://github.com/ddiu8081/chatgpt-demo#deploy-with-vercel)

2. Build chatgpt-demo locally:

    cd chatgpt-demo
    npm run build
    Run reserve proxy:
    cd dist

3. Run reserve proxy:


```javascript
# Make sure your binary file name is chatgpt-demo-proxy and already add to $PATH
chatgpt-demo-proxy --client-dir=./client --port=9999

# If you need a proxy to proxy /api/generate try:
# chatgpt-demo-proxy --client-dir=./client --port=9999 --proxy=http://127.0.0.0:1080
```

4. Visit [http://localhost:9999](http://localhost:9999) to access the deployed chatgpt-demo project.