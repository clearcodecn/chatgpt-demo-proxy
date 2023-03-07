#### Reserve proxy for chatgpt-demo projewct

If you don't know how to deploy `chatgpt-demo` locally and you want deploy `chatgpt-demo` at your server, you can try this project!

#### Implementation principle

![https://user-images.githubusercontent.com/34591322/223335950-8af037e8-d3aa-4ea5-ac3d-b5c281dcdc96.png](https://user-images.githubusercontent.com/34591322/223335950-8af037e8-d3aa-4ea5-ac3d-b5c281dcdc96.png)

#### Install 
If you have go environment, you can install it directly

    go get github.com/clearcodecn/chatgpt-demo-proxy

or you can download binaries at [release page](https://www.baidu.com)


#### How to use

1. [deploy `chatgpt-demo` by `Vercel`](https://github.com/ddiu8081/chatgpt-demo#deploy-with-vercel)

2. build `chatgpt-demo` locally  

```javascript
   cd chatgpt-demo 
   npm run build
```

3. run reserve proxy

```javascript
cd dist

# make sure your binary file name is chatgpt-demo-proxy and already add to $PATH
chatgpt-demo-proxy --client-dir=./client --port=9999 

# if you need a proxy to proxy /api/generate try:
# chatgpt-demo-proxy --client-dir=./client --port=9999 --proxy=http://127.0.0.0:1080 
```

4. visit [http://localhost:9999](http://localhost:9999)
