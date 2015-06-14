# mm3save

A library that tries to understand the Might and Magic 3 save game format.  And a set of utilities to perform common operations on the files.

## Why?

I wanted some nostaltic replay value from the role playing games of my youth, but without all the dungeon crawling grind at lower levels.

### Wait, isn't that cheating?

To me it's just playing a different game.  I get more enjoyment reverse engineering the file format than I do reloading my save game after being eaten by a pack of Moose Rat.

## Utility Programs

There are 2 useful utility programs.

### mm3print

This program reads your save game and prints out what we know.  I've tried to keep this updated as I add more features, but sometimes it lags behind, because I'm mostly modifying files and looking at them in game.

This is safe to use on any file, it opens the file read-only, and makes no modifications.

### mm3update

This program allows you to make modifications to the save game file.

**WARNING**: Don't be stupid, backup your save games before operating on them.  I'm not responsible for destroying your important save game, you are.

The easiest way to use this utility is to invoke it multiple times, each focused on a single thing.  For example, invoke it once to update the party properties (gold, gems, food).  Invoke it again to set a characters stats to 100.  Repeat this invocation once for each party member.  Then, if you want to modify items, invoke it once for each item slot you want to modify.  Isn't that a lot of invocations?  Yeah, but it keeps the command-line simple.

Examples:

Give party $1M gold, $1M gold in bank, 1M gems, 1M gems in bank, and 50k food.

```
$ mm3update -o SAVE01.MM3 -name "1a" -gold 1000000 -bankgold 1000000 -gems 1000000 -bankgems 1000000 -food 50000 SAVE01.MM3
setting save game name to '1a'
setting party food to 50000
setting party gold to 1000000
setting party gems to 1000000
setting bank gold to 1000000
setting bank gems to 1000000
```

Set Sir Cagegm's stats to 100, resistances to 100, and grant all skills.

```
$ mm3update -o SAVE01.MM3 -name "1a" -char "Sir Canegm" -mgt 100 -int 100 -per 100 -end 100 -spd 100 -acy 100 -lck 100 -fire 100 -cold 100 -elec 100 -poison 100 -energy 100 -magic 100 -skills "thievery,armsmaster,astrologer,bodybuilder,cartographer,crusader,directionsense,linguist,merchant,mountaineer,navigator,pathfinder,prayermaster,prestidigitator,swimmer,tracker,spotsecretdoors,dangersense" SAVE01.MM3
setting save game name to '1a'
updating character 'Sir Canegm'
setting MGT to 100
setting INT to 100
setting PER to 100
setting END to 100
setting SPD to 100
setting ACY to 100
setting LCK to 100
setting fire res to 100
setting cold res to 100
setting elec res to 100
setting poison res to 100
setting energy res to 100
setting magic res to 100
giving character the skills thievery,armsmaster,astrologer,bodybuilder,cartographer,crusader,directionsense,linguist,merchant,mountaineer,navigator,pathfinder,prayermaster,prestidigitator,swimmer,tracker,spotsecretdoors,dangersense
```

Repeat the same for other characters (output omitted)

```
$ mm3update -o SAVE01.MM3 -name "1a" -char "Crag Hack" -mgt 100 -int 100 -per 100 -end 100 -spd 100 -acy 100 -lck 100 -fire 100 -cold 100 -elec 100 -poison 100 -energy 100 -magic 100 -skills "thievery,armsmaster,astrologer,bodybuilder,cartographer,crusader,directionsense,linguist,merchant,mountaineer,navigator,pathfinder,prayermaster,prestidigitator,swimmer,tracker,spotsecretdoors,dangersense" SAVE01.MM3
$ mm3update -o SAVE01.MM3 -name "1a" -char "Maximus" -mgt 100 -int 100 -per 100 -end 100 -spd 100 -acy 100 -lck 100 -fire 100 -cold 100 -elec 100 -poison 100 -energy 100 -magic 100 -skills "thievery,armsmaster,astrologer,bodybuilder,cartographer,crusader,directionsense,linguist,merchant,mountaineer,navigator,pathfinder,prayermaster,prestidigitator,swimmer,tracker,spotsecretdoors,dangersense" SAVE01.MM3
$ mm3update -o SAVE01.MM3 -name "1a" -char "Dark Shade" -mgt 100 -int 100 -per 100 -end 100 -spd 100 -acy 100 -lck 100 -fire 100 -cold 100 -elec 100 -poison 100 -energy 100 -magic 100 -skills "thievery,armsmaster,astrologer,bodybuilder,cartographer,crusader,directionsense,linguist,merchant,mountaineer,navigator,pathfinder,prayermaster,prestidigitator,swimmer,tracker,spotsecretdoors,dangersense" SAVE01.MM3
$ mm3update -o SAVE01.MM3 -name "1a" -char "Resurectra" -mgt 100 -int 100 -per 100 -end 100 -spd 100 -acy 100 -lck 100 -fire 100 -cold 100 -elec 100 -poison 100 -energy 100 -magic 100 -skills "thievery,armsmaster,astrologer,bodybuilder,cartographer,crusader,directionsense,linguist,merchant,mountaineer,navigator,pathfinder,prayermaster,prestidigitator,swimmer,tracker,spotsecretdoors,dangersense" SAVE01.MM3
$ mm3update -o SAVE01.MM3 -name "1a" -char "Kastore" -mgt 100 -int 100 -per 100 -end 100 -spd 100 -acy 100 -lck 100 -fire 100 -cold 100 -elec 100 -poison 100 -energy 100 -magic 100 -skills "thievery,armsmaster,astrologer,bodybuilder,cartographer,crusader,directionsense,linguist,merchant,mountaineer,navigator,pathfinder,prayermaster,prestidigitator,swimmer,tracker,spotsecretdoors,dangersense" SAVE01.MM3
```

Replace Sir Canegm's first item slot with a `Fiery Obsidian Dragon Cutlass of Arrows`:

```
$ mm3update -o SAVE01.MM3 -name "1a" -char "Sir Canegm" -itemSlot 1 -itemDesc="fiery obsidian dragon cutlass of arrows" SAVE01.MM3
setting save game name to '1a'
updating character 'Sir Canegm'
setting item 1 to fiery obsidian dragon cutlass of arrows
```

Modifying an item slot leaves it in an unequipped state.  There are lots of combinations, to print the items tables invoke `mm3update --itemHelp`

## What else can I do?

- Change x,y position within current area
- Change direction
- Change HP/SP (temporary until rest)
- Change Exp (for very fast leveling)

## What can't I do?
- Change location (city, dungeon, outdoor region, etc)
- Change time (have some clues about this, but no motivation to pursue further)
- Toggle character status (dead/erradicated/stoned)?
- Grant awards (these are mapped in the library, but not exposed to utility yet)
- Toggle broken weapons (i seem to recall weapons breaking, and there is an unexplained field in the items, but toggling it to various values seemed to have no effect)
- Modify character main identity (race, class, gender, alignment) - we have these mapped in the library but not wired up to the utility
- Everything else...

## License

Apache License Version 2.0