# mp3RatingCleaner

mp3RatingCleaner moves songs to the trash that you have disliked or given a rating in iTunes.

This does depend on the `trash` command which can be installed with `brew install trash`. If you don't want to use the trash command, edit the code to just `rm` the file instead. However, the `trash` command allows you to get songs back if anything goes wrong by opening your trash and `putting back` items.

After running this program, you will have to go into iTunes, sort by rating, and highlight and delete all the 1 star songs (or disliked songs). This only removes the file from the file system and not iTunes.

## Usage

You first need to know where your iTunes library is located. This is typically `~/Music/iTunes/iTunes Library.xml`. This is not the actual library file, just an XML file that iTunes occasionally updates.

> If this program does not work, try opening and closing iTunes a few times to get it to render out this xml file.

To remove all 1 star songs, run

```
mp3RatingCleaner -itunes ~/Music/iTunes/iTunes\ Library.xml -rating 1
```

To remove all 2 star songs, run

```
mp3RatingCleaner -itunes ~/Music/iTunes/iTunes\ Library.xml -rating 2
```

To remove all disliked songs, run

```
mp3RatingCleaner -itunes ~/Music/iTunes/iTunes\ Library.xml -disliked
```

To remove all disliked songs and all 1 star songs, run

```
mp3RatingCleaner -itunes ~/Music/iTunes/iTunes\ Library.xml -rating 1 -disliked
```

You get the idea.
