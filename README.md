# DayOne

This tool parses the JSON file that DayOne export provides and writes one markdown file per day into a target directory

## Usage

```
dayone split input_file.csv existing_output_dir
```

## Format

I am ignoring the time zone information as I am only interested in the day (and 99% of the entries are in the same time zone)

```
{
  "metadata": {
    "version": "1.0"
  },
  "entries": [
    {
      "starred": false,
      "duration": 0,
      "creationDate": "1999-06-21T10:00:00Z",
      "isPinned": false,
      "modifiedDate": "2017-05-01T18:11:15Z",
      "timeZone": "Europe/Amsterdam",
      "text": "... \n\n ...",
      "isAllDay": false,
      "uuid": "24A355B7C88541A7AA4D599622FD5225"
    },
    ...
  ]
}
```
