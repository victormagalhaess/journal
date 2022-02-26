# Journal

Journal is a command line journal heavily inspired by [journal-cli](https://journalcli.app/).
It aims to keep all your entries in an easy to read and manipulate file, allowing you to take notes and consulting them later.
Entries saved on the journal are logged in the following way when read:
```sh
date: 26/02/2022
    entry: Why did you choose this day as an example? It is an ordinary day nothing special happened this day.
    hash: 2705810935
```
# Usage

To create an entry simply type:

```sh
journal add "Today I waste 15 minutes thinking in an entry example, :("
```

The entry will be saved as readable text in the cli installation folder. Entries that don't specifies a date will be saved as "today".
You can also specify a date using the add command:

```sh
journal add 14/2/2022 "Oh my god, I totally forgot to write about how normal 14/2/2022 was!"
```

```sh
journal add 26-02-2022 "In this day I thought that allowing people do use dashs as date separator may be a good idea!"
```

The add command also supports some special cases:

```sh
journal add today "I'm not thirsty at all today."
```

```sh
journal add yesterday "Wow, I drank a lot of water yesterday, in fact, I drank so much that I didn't had time to write an entry."
```

You can add as many entries per day as you want. When you try to read them the journal will log in the order they were wrote.

To read an entry, you must use the read command, I works pretty similar to the add command, so if you do not specify a date it will read the "today" entry:

```sh
journal read
```

You can also specify a date to consult:

```sh
journal read 14/2/2022
```

The read command also supports some special cases:

```sh
journal read today
```

```sh
journal read yesterday
```

Last but not least, you can also delete entries using the delete command, it follows the same principle as the other two flags, so if you do not specify a date the today entry will be deleted:

```sh
journal delete
```

You can also specify a date to delete:

```sh
journal delete 14/2/2022
```

The delete command also supports some special cases:

```sh
journal delete today
```

```sh
journal delete yesterday
```

Be aware that the delete command using days as filter will erase all the data from the selected day. The way to erase only a specific entry is to use the delete command passing the Hash of the entry:

 ```sh
journal delete 2705810935
```

# Build

journal is built in go, and can be built for your machine using:

```sh
make build
```
