gonote
====

A simple command line tool for saving quick notes. I wrote this
to remember new commands I learn for different programs.

#Getting Started

To install, run

    go get github.com/pscohn/gonote

gonote will save your notes in `~/.gonote.db`.

#Usage

Create a new category 'tmux'

    gonote -c tmux

List all categories

    gonote -l

Create a note in 'tmux' called 'tmux list-sessions'

    gonote -d tmux -n tmux list-sessions

Get all notes in 'tmux'

    gonote -g tmux

#Todo

- Command to delete category
- Command to delete notes
- Save date and sort by date
- Full test coverage
