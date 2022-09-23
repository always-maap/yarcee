```
arch=`uname -m`
kernel_path=$(pwd)"/hello-vmlinux.bin"
rootfs_path=$(pwd)"/hello-rootfs.ext4"
```

```
curl --unix-socket /tmp/firecracker.socket -i \                                                                                            INT ✘  system   21:21:23
      -X PUT 'http://localhost/boot-source'   \
      -H 'Accept: application/json'           \
      -H 'Content-Type: application/json'     \
      -d "{
            \"kernel_image_path\": \"${kernel_path}\",
            \"boot_args\": \"console=ttyS0 reboot=k panic=1 pci=off\"
       }"
```

```
curl --unix-socket /tmp/firecracker.socket -i \                                                                                            INT ✘  system   21:21:52
  -X PUT 'http://localhost/drives/rootfs' \
  -H 'Accept: application/json'           \
  -H 'Content-Type: application/json'     \
  -d "{
        \"drive_id\": \"rootfs\",
        \"path_on_host\": \"${rootfs_path}\",
        \"is_root_device\": true,
        \"is_read_only\": false
   }"
```

```
curl --unix-socket /tmp/firecracker.socket -i -X PUT 'http://localhost/actions' -H  'Accept: application/json' \                           INT ✘  system   21:22:08
-H  'Content-Type: application/json' -d'
{
   "action_type": "InstanceStart"
}'
```
