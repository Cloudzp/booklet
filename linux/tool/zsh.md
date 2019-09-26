# Oh-My-Zsh
> [install zsh] https://github.com/robbyrussell/oh-my-zsh/wiki/Installing-ZSH 
>
> [install oh-my-zsh ]https://github.com/robbyrussell/oh-my-zsh
>
> [install plugin autosuggestions]https://github.com/zsh-users/zsh-autosuggestions/blob/master/INSTALL.md

## Install zsh and oh-my-zsh
```
$ apt install zsh
$ sh -c "$(wget -O- https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
```
## Change Themes
Once you find a theme that you'd like to use, you will need to edit the ~/.zshrc file. You'll see an environment variable (all caps) in there that looks like:

`ZSH_THEME="ys"`

## Install plugin
### 1.Clone this repository into $ZSH_CUSTOM/plugins (by default ~/.oh-my-zsh/custom/plugins)
````
$ git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
````
### 2.Add the plugin to the list of plugins for Oh My Zsh to load (inside ~/.zshrc):

`plugins=(zsh-autosuggestions)`

### 3.Start a new terminal session.