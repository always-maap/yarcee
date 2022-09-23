dd if=/dev/zero of=rootfs.ext4 bs=1M count=1000
mkfs.ext4 rootfs.ext4
mkdir -p /tmp/my-rootfs
mount rootfs.ext4 /tmp/my-rootfs

docker run -i --rm \
    -v /tmp/my-rootfs:/my-rootfs \
    -v "$(pwd)/fire:/usr/local/bin/fire" \
    -v "$(pwd)/openrc.sh:/etc/init.d/fire" \
    --net=host \
    alpine sh < setup-alpine.sh

umount /tmp/my-rootfs

# rootfs available under `rootfs.ext4`