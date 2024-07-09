# Generative Art Experiments

(Algorithmic art, not generative AI models.)

## Bootstrap

This repo uses [hermit](https://cashapp.github.io/hermit/), so it
should install the right versions of things automatically. Let me know
if anything doesn't work.

**Note:** unfortunately, hermit does not (currently) work on Windows,
so there you'll need to create your own Python virtual environment,
and install any other tools.

```
git clone https://github.com/zellyn/genart
cd genart
. ./bin/activate-hermit
pip install -r requirements.txt
```

## Projects

### [Schotter](./research/schotter)

Experiments related to Georg Nees' "Schotter" ("Gravel") aka "Cubic
Disarray".

### [23-ecke](./research/23-ecke)

Experiments related to Georg Nees' "23 ecke" ("23 corners") aka
"Polygon of 23 vertices".
