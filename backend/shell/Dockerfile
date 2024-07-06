FROM alpine:latest

RUN apk update && \
    apk add --no-cache \
        bash \
        sudo \
        openssh \
        curl \
        wget \
        jq \
        vim \
        nano \
        tmux \
        git \
        python3 \
        nodejs \
        npm \
        gcc \
        libc-dev \
        make \
        openssl-dev \
        ca-certificates


# user environment
ARG USERNAME=cloudshell-user
ARG USER_UID=1000
ARG USER_GID=1000

# user setup
RUN addgroup -g $USER_GID $USERNAME && \
    adduser -D -u $USER_UID -G $USERNAME -h /home/$USERNAME -s /bin/bash $USERNAME && \
    echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME && \
    mkdir -p /home/$USERNAME/.ssh && \
    chown -R $USERNAME:$USERNAME /home/$USERNAME

# ssh setup
RUN sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config && \
    sed -i 's/#PermitEmptyPasswords no/PermitEmptyPasswords yes/' /etc/ssh/sshd_config && \
    sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin no/' /etc/ssh/sshd_config


# Expose SSH port
EXPOSE 22

# Start SSH daemon
CMD ["/usr/sbin/sshd", "-D"]
