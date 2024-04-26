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
0[Dreadnaught IPA] --> 1[Hercules Double IPA]
0[Dreadnaught IPA] --> 2(Bell’s Expedition)
1[Hercules Double IPA] --> 2(Bell’s Expedition)
1[Hercules Double IPA] --> 3((Old Rasputin Russian Imperial Stout))
2(Bell’s Expedition) --> 3((Old Rasputin Russian Imperial Stout))
2(Bell’s Expedition) --> 4[Weihenstephaner Hefeweissbier]
3((Old Rasputin Russian Imperial Stout)) --> 4[Weihenstephaner Hefeweissbier]
3((Old Rasputin Russian Imperial Stout)) --> 5{Founders Kentucky Breakfast}
4[Weihenstephaner Hefeweissbier] --> 5{Founders Kentucky Breakfast}
4[Weihenstephaner Hefeweissbier] --> 6((Nugget Nectar))
5{Founders Kentucky Breakfast} --> 6((Nugget Nectar))
5{Founders Kentucky Breakfast} --> 7[Hennepin]
6((Nugget Nectar)) --> 7[Hennepin]
6((Nugget Nectar)) --> 8{Edmund Fitzgerald Porter}
7[Hennepin] --> 8{Edmund Fitzgerald Porter}
7[Hennepin] --> 9((Hop Rod Rye))
8{Edmund Fitzgerald Porter} --> 9((Hop Rod Rye))
8{Edmund Fitzgerald Porter} --> 10{Edmund Fitzgerald Porter}
9((Hop Rod Rye)) --> 10{Edmund Fitzgerald Porter}
9((Hop Rod Rye)) --> 11[Oak Aged Yeti Imperial Stout]
10{Edmund Fitzgerald Porter} --> 11[Oak Aged Yeti Imperial Stout]
10{Edmund Fitzgerald Porter} --> 12([Maharaj])
11[Oak Aged Yeti Imperial Stout] --> 12([Maharaj])
11[Oak Aged Yeti Imperial Stout] --> 13[Hop Rod Rye]
12([Maharaj]) --> 13[Hop Rod Rye]
12([Maharaj]) --> 14[HopSlam Ale]
13[Hop Rod Rye] --> 14[HopSlam Ale]
13[Hop Rod Rye] --> 15((Stone Imperial Russian Stout))
14[HopSlam Ale] --> 15((Stone Imperial Russian Stout))
14[HopSlam Ale] --> 16[Chocolate St]
15((Stone Imperial Russian Stout)) --> 16[Chocolate St]
15((Stone Imperial Russian Stout)) --> 17[Double Bastard Ale]
16[Chocolate St] --> 17[Double Bastard Ale]
16[Chocolate St] --> 18{Edmund Fitzgerald Porter}
17[Double Bastard Ale] --> 18{Edmund Fitzgerald Porter}
17[Double Bastard Ale] --> 19(Oak Aged Yeti Imperial Stout)
18{Edmund Fitzgerald Porter} --> 19(Oak Aged Yeti Imperial Stout)
18{Edmund Fitzgerald Porter} --> 20[Hercules Double IPA]
19(Oak Aged Yeti Imperial Stout) --> 20[Hercules Double IPA]
19(Oak Aged Yeti Imperial Stout) --> 21[Sierra Nevada Celebration Ale]
20[Hercules Double IPA] --> 21[Sierra Nevada Celebration Ale]
20[Hercules Double IPA] --> 22[Schneider Aventinus]
21[Sierra Nevada Celebration Ale] --> 22[Schneider Aventinus]
21[Sierra Nevada Celebration Ale] --> 23{Hennepin}
22[Schneider Aventinus] --> 23{Hennepin}
22[Schneider Aventinus] --> 24[Hennepin]
23{Hennepin} --> 24[Hennepin]
23{Hennepin} --> 25([Celebrator Doppelbock])
24[Hennepin] --> 25([Celebrator Doppelbock])
24[Hennepin] --> 26{Westmalle Trappist Tripel}
25([Celebrator Doppelbock]) --> 26{Westmalle Trappist Tripel}
25([Celebrator Doppelbock]) --> 27{Two Hearted Ale}
26{Westmalle Trappist Tripel} --> 27{Two Hearted Ale}
26{Westmalle Trappist Tripel} --> 28{Shakespeare Oatmeal}
27{Two Hearted Ale} --> 28{Shakespeare Oatmeal}
{{< /mermaid >}}
