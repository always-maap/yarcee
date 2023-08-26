<div align="center">
  <a href="">
    <img height="150px;" src="https://raw.githubusercontent.com/always-maap/yarcee/master/client/public/yarcee.png" alt="ts-collection logo" />
  </a>
  
  <p>Yet Another Remote Code Execution Engine</p>

</div>

## Overview

YARCEE(yet another remote code execution engine) is a code running service that relies on
Firecracker to spawn microVMs and execute the code over HTTP.

## Services

| Service                    | Description                                                                    |
| :------------------------- | ------------------------------------------------------------------------------ |
| [`api`](./api)             | REST API service using Go lang and postgresql                                  |
| [`client`](./client)       | Client app using Next.js app and TypeScript                                    |
| [`frontline`](./frontline) | API service that runs inside the VM and spawn child_process written in Go lang |
| [`vmvisor`](./vmvisor)     | VM manager to spawn firecracker VM using Go lang and firecracker-sdk           |

## Contributing

We're open to all community contributions! If you'd like to contribute in any way, please first read our Contributing Guide.

## resources:

- https://github.com/firecracker-microvm/firecracker/blob/main/docs/getting-started.md
- https://arun-gupta.github.io/firecracker-getting-started/
- https://stanislas.blog/2021/08/firecracker/
- https://jvns.ca/blog/2021/01/23/firecracker--start-a-vm-in-less-than-a-second/

## License

under [MIT-licensed](./LICENSE).
