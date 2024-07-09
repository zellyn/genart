# Georg Nees random notes

Best picture of 23-ecke (Polygon of 23 vertices)
http://evanrosica.com/GANArt/Final%20Paper_Compressed.pdf

Description of "Computer-Grafik" (Nees 1965)
http://dada.compart-bremen.de/item/exhibition/325

Description of "rot 19"
http://dada.compart-bremen.de/item/publication/362

Rot 19 pdf [Rot_19_Computer-Grafik_1968](./Rot_19_Computer-Grafik_1968.pdf)
source: https://monoskop.org/images/b/b1/Rot_19_Computer-Grafik_1968.pdf

Exhibition about Rot ("Red")
https://zkm.de/en/exhibition/2024/07/reading-red
https://zkm.de/en/project/reading-rot-again

Brief bio
https://monoskop.org/Georg_Nees

PDF with image of part of printer's draft/proof of Rot-19 - 23-ecke
http://www.generativeart.com/GA2019_web/22_Ga%C3%ABtan%20Robillard_Alain%20Lioret_paper_168x240.pdf

Diaphenes Rot 19 (€3)
https://www.diaphanes.net/buch/artikel/1928

Rot 19 for sale ($375)
https://www.betweenthecovers.com/pages/books/553295/computer-grafik


Nake personal recollection
Has prose descriptions of the 1965 gallery exhibition and reactions
https://www.pismin.com/10.1145/1056224.1056234
doi:10.1145/1056224.1056234

What is generative art?
Mentions Nees and the 1965 exhibit
https://creativecoding.soe.ucsc.edu/courses/cmpm202_w20/texts/Boden_Edmonds_WhatIsGenerativeArt.pdf

Earlier print of 23-ecke
https://archive.org/details/grkg-1960-bd.-1-h-1-h-4/1964/GRKG-1964-Bd5-H3%264/page/123/mode/2up

* 23-ecke

1d actual E/W
=============

2223 Winning seed: 1034132409   (errs=1)
All winning seeds:
 1034132409

1d actual N/S
=============

223 Winning seed: 367055537   (errs=1)
3 Winning seed: 1582835865   (errs=1)
All winning seeds:
 367055537
 1582835865

1d actual reversed E/W
======================

23 Winning seed: 1264012749   (errs=1)
22223 Winning seed: 1365575867   (errs=1)
All winning seeds:
 1264012749
 1365575867

1d actual reversed N/S
======================
2222223 Winning seed: 657107221   (errs=1)
All winning seeds:
 657107221


# Communication log

## Mon Jul  1 10:07:07 EDT 2024 - gaetanrobillard.studio@gmail.com

Hi there,

I've been researching some of Georg Nees' early work, attempting to recover the specific random seeds he used, to recreate the artwork exactly. You can see an example of the process and result documented at https://zellyn.com/2024/06/schotter-1/ and https://zellyn.com/2024/06/schotter-2/.

I ran across your paper, "A Vision without a Sight: From Max Bense's Theory to the Dialectic of Programmed Images", while looking for high-resolution images of the "23 ecke" artwork, and noticed Figure 2: "Georg Nees, 23-Edge: detailed view from editor’s draft, Rot n° 19, 1965."

The image in your paper offers the tantalizing possibility of getting a high-resolution image of the printer's proof! So far the best image I've found online is in a 1964 edition of Grundlagenstudien aus Kybernetik und Geisteswissenschaft (GRKG) on archive.org

Do you have an image of the full 23-ecke proof image that you would be able to share? Or, failing that, I'm curious where you got it: did you visit ZKM in person, or find it online? I'd appreciate any pointers you have on tracking down a high-resolution image of it.

Thank you,

Zellyn Hunter

## Jun 24, 2024, 11:09 PM - errosica@gmail.com

Hi Mr. Rosica,

I have a somewhat odd question about your "Combinatoric emergence in GANs and style transfer networks" paper:

What is the source of the image labeled "Figure 5.1. Georg Nees: 23-Ecke (Polygon of 23 vertices) - 1965"?

I've been tracking down source code and trying to exactly recreate Georg Nees' artwork — you can see my work so far in the last two posts on zellyn.com. In the course of trying to recreate his "23-Ecke" work, I ran across your paper, which includes the highest-resolution image of the work I have been able to find yet!

I was curious if you remember where that image came from — I'd love to learn as much about it as possible.

Thanks!

Zellyn Hunter

### Reply:

Hello, and thanks for the email.

I really like your idea to recreate these "classic" computer art images. One note on your code in "part 1" is that if you just want to rename a function from a library to X, you can just import it as X. E.g. the following are equivalent:

def SIN(x):
    return math.sin(x)

from math import sin as SIN

Unfortunately, I cannot find the source of the image. I tried looking through the sources in the bibliography, and using google reverse image search without any luck. It may be that the original source is no longer available.

For research purposes, I've attached the exact image I used in my paper below.  It's saved as "Georg-Nees-23-Ecke-1965.PNG" with resolution 1170 × 1602, and I downloaded it on Dec 30, 2018. Please let me know if you find the image source so I can properly cite it.

Best,

Evan Rosica
Georg-Nees-23-Ecke-1965.PNG

# Mon, Jul 1, 2024 at 10:25 AM - ZKM

From:	Zellyn Hunter <zellyn@gmail.com>
To:	sammlung-und-archive@zkm.de, bildanfragen@zkm.de
Subject:	Question about the printer's draft for the rot 19 journal

Hello,

I'm on a project seeking to recreate some of Georg Nees' early artwork, by
translating his original code, and also searching for the initial seeds he
used in his random number routines. You can see an example of my process
and the results documented at https://zellyn.com/2024/06/schotter-1/ and
https://zellyn.com/2024/06/schotter-2/.

I'm currently trying to recreate his "23 ecke" (also sometimes known as
"Polygons of 23 vertices"), and I ran across a tantalizing lower-resolution
fragment of an artifact in your collection: it is referenced in the
bibliography of the paper as "Fig2. Georg Nees, 23-Edge: detailed view from
editor=E2=80=99s draft, Rot n=C2=B0 19, 1965. ZKM | Center for Art and Medi=
a Karlsruhe /
Elisabeth Walther-Bense Estate / ZKM-01".  I believe that is the collection
referenced online at https://zkm.de/en/elisabeth-walther-bense

The paper in question is "A Vision without a Sight: From Max Bense's Theory
to the Dialectic of Programmed Images", by Ga=C3=ABtan Robillard and Prof. =
Alain
Lioret, available at
http://www.generativeart.com/GA2019_web/22_Ga%C3%ABtan%20Robillard_Alain%20=Lioret_paper_168x240.pdf
on the website of the XXII Generative Art Conference in 2019.
<https://generativeart.com/GA2019_web/Generative%20Art%202019%20index.htm>

Here is the image I'm interested in:
[image: image.png]

I was wondering if it is possible to get a high-resolution scan of that
printer's draft? So far the best image I have been able to find online of
"23 ecke" is in a 1964 edition of Grundlagenstudien aus Kybernetik und
Geisteswissenschaft (GRKG) on archive.org
<https://archive.org/details/grkg-1960-bd.-1-h-1-h-4/1964/GRKG-1964-Bd5-H3%=
264/page/123/mode/2up>
.

Thank you for your time,

Zellyn Hunter
