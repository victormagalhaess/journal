# Journal

Journal is a command line journal heavily inspired by [journal-cli](https://journalcli.app/).
It aims to keep all your entries in a easy to read and manipulate file, allowing you to take notes and consulting them lates.

# Usage

To create an entry simply type:

```sh
journal add "Today I waste 15 minutes thinking in an entry example, :("
```

The entry will be saved as readable text in the cli installation folder. Entries that don't specifies a date will be saved as "today".
You can also specify a date using the add command: (Be aware that it is not allowed to write entries about days in the future).

```sh
journal add 14/2/2022 "Oh my god, I totally forgot to write about how normal 14/2/2022 was!"
```

The add command also supports some special cases:

```sh
journal add today "I'm not thirsty at all today."
```

```sh
journal add yesterday "Wow, I drank a lot of water yesterday, in fact, I drank so much that I didn't had time to write an entry."
```

If you try to write two entries in the same date, the entries will be appended in one big entry.

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

# Build

journal is built in go, and can be built for your machine using:

```sh
make build
```
