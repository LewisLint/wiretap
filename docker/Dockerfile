FROM archlinux:latest

# Step 1: Initialize keys and install core packages (Cached)
RUN pacman-key --init \
    && pacman-key --populate archlinux \
    && pacman -Syu --noconfirm \
    && pacman -S --needed --noconfirm base-devel go git shadow sudo zsh \
    && pacman -Scc --noconfirm

# Step 2: Set up user environment permissions
RUN useradd --create-home --uid 1000 --user-group --shell /usr/bin/zsh vscode \
    && printf 'vscode ALL=(ALL) NOPASSWD: ALL\n' > /etc/sudoers.d/vscode \
    && chmod 0440 /etc/sudoers.d/vscode

# Step 3: Clone dotfiles and fix ownership permissions
RUN git clone https://github.com/lewislint/dotfiles /home/vscode/dots \
    && cp /home/vscode/dots/rose/.zshrc /home/vscode/.zshrc \
    && cp /home/vscode/dots/rose/.vimrc /home/vscode/.vimrc \
    && chown -R vscode:vscode /home/vscode

USER vscode
WORKDIR /home/vscode

