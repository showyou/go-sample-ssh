FROM debian:latest

RUN apt-get update && apt-get install -y openssh-server
RUN mkdir /var/run/sshd
RUN echo "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICjUwBLfg3fpt3XOGPBbSJpeAPnkRp/f/i6WtlllU7s+ yuki@MacBook-Air.local" > /root/.ssh/authorized_keys

EXPOSE 22

COPY id_ed25519_for_docker_linux_ssh.pub /root/.ssh/authorized_keys
CMD chmod 600 /root/.ssh/authorized_keys && \
    /usr/sbin/sshd -D
