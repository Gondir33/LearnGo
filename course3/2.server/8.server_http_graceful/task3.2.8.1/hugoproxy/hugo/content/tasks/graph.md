---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

# Построение графа

{{< mermaid >}}
graph LR
0[Weihenstephaner Hefeweissbier] --> 1((90 Minute IPA))
0[Weihenstephaner Hefeweissbier] --> 2[Two Hearted Ale]
1((90 Minute IPA)) --> 2[Two Hearted Ale]
1((90 Minute IPA)) --> 3[Maharaj]
2[Two Hearted Ale] --> 3[Maharaj]
2[Two Hearted Ale] --> 4([Dreadnaught IPA])
3[Maharaj] --> 4([Dreadnaught IPA])
3[Maharaj] --> 5([Oak Aged Yeti Imperial Stout])
4([Dreadnaught IPA]) --> 5([Oak Aged Yeti Imperial Stout])
4([Dreadnaught IPA]) --> 6{Maharaj}
5([Oak Aged Yeti Imperial Stout]) --> 6{Maharaj}
5([Oak Aged Yeti Imperial Stout]) --> 7(Schneider Aventinus)
6{Maharaj} --> 7(Schneider Aventinus)
6{Maharaj} --> 8{Sierra Nevada Celebration Ale}
7(Schneider Aventinus) --> 8{Sierra Nevada Celebration Ale}
7(Schneider Aventinus) --> 9[Maudite]
8{Sierra Nevada Celebration Ale} --> 9[Maudite]
8{Sierra Nevada Celebration Ale} --> 10{Alpha King Pale Ale}
9[Maudite] --> 10{Alpha King Pale Ale}
9[Maudite] --> 11([Trappistes Rochefort 10])
10{Alpha King Pale Ale} --> 11([Trappistes Rochefort 10])
10{Alpha King Pale Ale} --> 12{Hennepin}
11([Trappistes Rochefort 10]) --> 12{Hennepin}
11([Trappistes Rochefort 10]) --> 13(Old Rasputin Russian Imperial Stout)
12{Hennepin} --> 13(Old Rasputin Russian Imperial Stout)
12{Hennepin} --> 14[Stone Imperial Russian Stout]
13(Old Rasputin Russian Imperial Stout) --> 14[Stone Imperial Russian Stout]
13(Old Rasputin Russian Imperial Stout) --> 15[Sierra Nevada Bigfoot Barleywine Style Ale]
14[Stone Imperial Russian Stout] --> 15[Sierra Nevada Bigfoot Barleywine Style Ale]
14[Stone Imperial Russian Stout] --> 16([Old Rasputin Russian Imperial Stout])
15[Sierra Nevada Bigfoot Barleywine Style Ale] --> 16([Old Rasputin Russian Imperial Stout])
15[Sierra Nevada Bigfoot Barleywine Style Ale] --> 17[Double Bastard Ale]
16([Old Rasputin Russian Imperial Stout]) --> 17[Double Bastard Ale]
16([Old Rasputin Russian Imperial Stout]) --> 18(Samuel Smith’s Oatmeal Stout)
17[Double Bastard Ale] --> 18(Samuel Smith’s Oatmeal Stout)
17[Double Bastard Ale] --> 19{Ruination IPA}
18(Samuel Smith’s Oatmeal Stout) --> 19{Ruination IPA}
18(Samuel Smith’s Oatmeal Stout) --> 20[Celebrator Doppelbock]
19{Ruination IPA} --> 20[Celebrator Doppelbock]
19{Ruination IPA} --> 21([Alpha King Pale Ale])
20[Celebrator Doppelbock] --> 21([Alpha King Pale Ale])
20[Celebrator Doppelbock] --> 22([Founders Kentucky Breakfast])
21([Alpha King Pale Ale]) --> 22([Founders Kentucky Breakfast])
21([Alpha King Pale Ale]) --> 23[Yeti Imperial Stout]
22([Founders Kentucky Breakfast]) --> 23[Yeti Imperial Stout]
{{< /mermaid >}}
