# Journal

Journal is a command line journal heavily inspired by [journal-cli](https://journalcli.app/).
It aims to keep all your entries in a easy to read and manipulate file, allowing you to take notes and consulting them lates.

# Usage

To create an entry simply type:

```sh
journal "Today I waste 15 minutes thinking in an entry example, :("
```

The entry will be saved as readable text in the cli installation folder. Entries that don't specifies a date will be saved as "today".
You can also specify a date using the -n flag: (Be aware that it is not allowed to write entries about days in the future).

```sh
journal -n 14/2/2022 "Oh my god, I totally forgot to write about how normal 14/2/2022 was!"
```

The -n flag also supports some special cases:

```sh
journal -n today "I'm not thirsty at all today."
```

```sh
journal -n yesterday "Wow, I drank a lot of water yesterday, in fact I drank so much that I didn't had time to write an entry."
```

If you try to write two entries in the same date, the entries will be appended in one big entry.

To read an entry, you must use the -r flag, I works pretty similar to the -n flag, so if you do not specify a date it will read the "today" entry:

```sh
journal -r
```

You can also specify a date to consult:

```sh
journal -r 14/2/2022
```

The -r flag also supports some special cases:

```sh
journal -r today
```

```sh
journal -r yesterday
```

Last but not least, you can also delete entries using the flag -d, it follows the same principle as the other two flags, so if you do not specify a date the today entry will be deleted:

```sh
journal -d
```

You can also specify a date to delete:

```sh
journal -d 14/2/2022
```

The -d flag also supports some special cases:

```sh
journal -d today
```

```sh
journal -d yesterday
```

# Build

journal is built in go, and can be built for your machine using:

```sh
make build
```
