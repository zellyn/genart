import marimo

__generated_with = "0.6.24"
app = marimo.App(width="medium")


@app.cell(hide_code=True)
def __(mo):
    mo.md("# Experiments recreating “23-ecke” by Georg Nees")
    return


@app.cell
def __():
    import io
    import math
    import random
    import drawsvg as draw
    import marimo as mo
    from PIL import Image
    from datetime import datetime
    return Image, datetime, draw, io, math, mo, random


@app.cell
def __():
    class Random:

        def __init__(self, seed):
            self.JI = seed

        def next(self, JA, JE):
            self.JI = self.JI * 5 % 2147483648
            return self.JI / 2147483648 * (JE - JA) + JA
    return Random,


@app.cell
def __():
    JS1 = 1306859721
    JS2 = 1485627309
    JS3 = 1649173265
    JS4 = 1805297143
    JS5 = 1973195467
    JS6 = 2013911545
    JS7 = 1922110153
    JS8 = 1769133315

    SEEDS = [JS1, JS2, JS3, JS4, JS5, JS6, JS7, JS8]
    return JS1, JS2, JS3, JS4, JS5, JS6, JS7, JS8, SEEDS


@app.cell(hide_code=True)
def __(mo):
    mo.md("# 23-ecke")
    return


@app.cell(hide_code=True)
def __(SEEDS, mo):
    ew_seeds = SEEDS
    ns_seeds = SEEDS

    seed1_slider = mo.ui.slider(1, len(ew_seeds), value=2)
    seed2_slider = mo.ui.slider(1, len(ns_seeds), value=5)

    # [seed1_slider, seed2_slider]
    mo.md(f'''
    Horizontal seed index (1–{len(ew_seeds)}): {seed1_slider}<br/>
    Vertical seed index (1–{len(ns_seeds)}): &nbsp; &nbsp; {seed2_slider}
    ''')
    return ew_seeds, ns_seeds, seed1_slider, seed2_slider


@app.cell(hide_code=True)
def __(Random, draw, ew_seeds, mo, ns_seeds, seed1_slider, seed2_slider):
    def draw_figures2(QUER, HOCH, XMAL, YMAL, inset, N, S1, S2, diag=False, stroke_width=0.3, svg_width=160, svg_height=200, swizzle=True, generators=2):
        if swizzle:
            rotate = 180
        else:
            rotate = 0

        def L(p, x, y):
            if swizzle:
                p.L(y, x)
            else:
                p.L(x, -y)
        def M(p, x, y):
            if swizzle:
                p.M(y, x)
            else:
                p.M(x, -y)

        JLI = inset / 2
        JRE = QUER - inset / 2
        JUN = inset / 2
        JOB = HOCH - inset / 2
        x = y = 0

        r1 = Random(S1)
        r2 = Random(S2)
        if generators == 1:
            r2 = r1

        d = draw.Drawing(svg_width, svg_height, origin='center', style="background-color:#eae6e2")
        g = draw.Group(transform=f"rotate({rotate})", stroke='#41403a', stroke_width=f'{stroke_width}', fill='none', stroke_linecap="round", stroke_linejoin="round")
        p = [None]

        def serie(QUER, HOCH, XMAL, YMAL, FIGUR):
          P = -QUER * XMAL * 0.5
          Q = YANF = -HOCH * YMAL * 0.5

          for COUNTX in range(1, XMAL+1):
            Q = YANF
            for COUNTY in range(1, YMAL+1):
              p[0] = draw.Path()
              FIGUR(P, Q, p[0])
              g.append(p[0])
              Q = Q + HOCH
            P = P + QUER

        def elirr(P, Q, p):
            JA1 = P + JLI
            JE1 = P + JRE
            JA2 = Q + JUN
            JE2 = Q + JOB
            P1 = r1.next(JA1, JE1)
            Q1 = r2.next(JA2, JE2)
            M(p, P1, Q1)
            X, Y = P1, Q1
            for _ in range(1, N+1):
                X = r1.next(JA1, JE1)
                L(p, X, Y)
                Y = r2.next(JA2, JE2)
                L(p, X, Y)
            if diag:
                L(p, P1, Q1)
            else:
                L(p, P1, Y)
                X = P1
                L(p, X, Q1)

        serie(QUER, HOCH, XMAL, YMAL, elirr)

        d.append(g)
        return d

    _seed1 = ew_seeds[seed1_slider.value-1]
    _seed2 = ns_seeds[seed2_slider.value-1]
    svg = draw_figures2(20, 20, 8, 6, 0, 20, _seed1, _seed2, diag=False, stroke_width=0.3, svg_width=140, svg_height=180, swizzle=True, generators=2).set_render_size(w=800).as_svg()
    mo.Html(svg)
    return draw_figures2, svg


@app.cell
def __(Image, mo):
    ecke_image1 = Image.open('23-ecke.png')
    ecke_image2 = Image.open('23-ecke-gc.png')
    subimage_row_get, subimage_row_set = mo.state(1)
    subimage_col_get, subimage_col_set = mo.state(1)
    comment_get, comment_set = mo.state('')
    return (
        comment_get,
        comment_set,
        ecke_image1,
        ecke_image2,
        subimage_col_get,
        subimage_col_set,
        subimage_row_get,
        subimage_row_set,
    )


@app.cell
def __(
    mo,
    subimage_col_get,
    subimage_col_set,
    subimage_row_get,
    subimage_row_set,
):
    im_slider_x = mo.ui.slider(1, 20, show_value=True, value=subimage_col_get(), on_change=subimage_col_set)
    im_slider_y = mo.ui.slider(1, 14, show_value=True, value=subimage_row_get(), on_change=subimage_row_set)
    return im_slider_x, im_slider_y


@app.cell(hide_code=True)
def __(im_slider_x, im_slider_y, mo):
    mo.md(
        f"""
        Horizontal sub-image: {im_slider_x}<br/>
        Vertical sub-image: &nbsp; &nbsp; &nbsp;   {im_slider_y}<br/>
        """
    )
    return


@app.cell(hide_code=True)
def __(
    comment_get,
    ecke_image1,
    ecke_image2,
    io,
    mo,
    subimage_col_get,
    subimage_row_get,
):
    def _map_comment(comment):
        mapping = {
            '/' : '/',
            '\\' : '\\',
            'n': 'narrow',
            's': 'short',
            't': 'tall',
            'w': 'wide',
            'd': 'down',
            'u': 'up',
            '.': '',
        }
        res = []
        for c in comment:
            res.append(mapping[c])
        return '<br/>'.join(res)

    _x1_sub_size = ecke_image1.width / 20
    _y1_sub_size = ecke_image1.height / 14
    _x2_sub_size = ecke_image2.width / 20
    _y2_sub_size = ecke_image2.height / 14

    _col = subimage_col_get()
    _row = subimage_row_get()
    _x = _col - 1
    _y = 14-_row
    _crop1 = (
        _x * _x1_sub_size,
        _y * _y1_sub_size,
        (_x + 1) * _x1_sub_size,
        (_y + 1) * _y1_sub_size,
        )
    _crop2 = (
        _x * _x2_sub_size,
        _y * _y2_sub_size,
        (_x + 1) * _x2_sub_size,
        (_y + 1) * _y2_sub_size,
        )
    _im1 = ecke_image1.crop(_crop1)
    _f1 = io.BytesIO()
    _im1.save(_f1, format='png')
    _im2 = ecke_image2.crop(_crop2)
    _f2 = io.BytesIO()
    _im2.save(_f2, format='png')
    mo.hstack(items=[
    mo.image(src=_f1, width=600, style={'image-rendering': 'pixelated', 'border': '1px solid black'}),
    mo.image(src=_f2, width=600, style={'image-rendering': 'pixelated', 'border': '1px solid black'}),
        mo.md(f'''
    <span style='font-size:3em'><b>Comments for ({_col},{_row})</b></span>

    <span style='font-size:2em'>{_map_comment(comment_get())}</span>
    ''')
    ],
              align='center',
              justify='start'
    )
    return


@app.cell(hide_code=True)
def __(
    comment,
    comment_set,
    comments,
    mo,
    subimage_col_get,
    subimage_col_set,
    subimage_row_get,
    subimage_row_set,
):
    def _click(dir):
        row, col = subimage_row_get(), subimage_col_get()
        if dir == 'next':
            row += 1
            if row > 14:
                row = 1
                col += 1
                if col > 20:
                    col = 1
                subimage_col_set(col)
        else:
            row -= 1
            if row < 1:
                row = 14
                col -= 1
                if col < 1:
                    col = 20
                subimage_col_set(col)
        subimage_row_set(row)
        comment_set(comments[row-1][col-1])

    def _classify(val):
        row, col = subimage_row_get()-1, subimage_col_get()-1
        global comment
        current = comments[row][col]
        if current == '':
            current = '....'
        if len(current) == 3:
            current += '.'
        old = [x for x in current]

        if val in ['/', '\\', '.clear_slope']:
            old[0] = val[0]
        elif val in ['narrow', 'wide', '.clear_width']:
            old[1] = val[0]
        elif val in ['short', 'tall', '.clear_height']:
            old[2] = val[0]
        elif val in ['up', 'down', '.clear_dir']:
            old[3] = val[0]
        else:
            raise Exception(f"unknown classification:'{val}'")
        new_val = ''.join(old)
        if new_val == '....':
            new_val = ''
        comments[row][col] = new_val
        comment_set(new_val)

    buttons = mo.ui.array([
        mo.ui.button(label="prev", value='prev', on_change=_click),
        mo.ui.button(label="next", value='next', on_change=_click),
        mo.ui.button(label='/', value='/', on_change=_classify),
        mo.ui.button(label='\\', value='\\', on_change=_classify),
        mo.ui.button(label='narrow', value='narrow', on_change=_classify),
        mo.ui.button(label='wide', value='wide', on_change=_classify),
        mo.ui.button(label='short', value='short', on_change=_classify),
        mo.ui.button(label='tall', value='tall', on_change=_classify),
        mo.ui.button(label='up', value='up', on_change=_classify),
        mo.ui.button(label='down', value='down', on_change=_classify),
        mo.ui.button(label='clear width', value='.clear_width', on_change=_classify),
        mo.ui.button(label='clear height', value='.clear_height', on_change=_classify),
        mo.ui.button(label='clear slope', value='.clear_slope', on_change=_classify),
        mo.ui.button(label='clear direction', value='.clear_dir', on_change=_classify),
    ])
    mo.hstack(justify='start', items=buttons)
    return buttons,


@app.cell
def __(comment_get, comments):
    print('comments = ' + '\n'.join(['['] + [f'  {str(cs)},' for cs in comments] + [']']) + comment_get()[:0])
    return


@app.cell
def __():
    comments = [
      ['/nsd', '\\nsd', '\\nsu', '/n.u', '/nsu', '\\wsu', '\\ns', '', '/wsu', '\\.sd', '', '\\nsu', '/n.', '/.su', '\\ntu', '\\ns', '\\nsu', '\\n.d', '/ns', ''],
      ['', '', '', '\\n.', '/nt', '/wsd', '/ns', '', '', '/ns', '/nsu', '', '/ntd', '/ns', '/nsu', '', '/wsu', '/nsu', '', ''],
      ['/w.u', '/n.', '\\ns', '', '/..', '/nsd', '\\ws', '/.sd', '\\n.d', '/.s', '\\nsd', '', '/.s', '/.td', '/..d', '\\ns', '/n.d', '\\.s', '/.s', ''],
      ['/nsu', '', '/n.u', '/ns', '/.su', '/n.', '/n.', '\\n.', '\\.sd', '/nsu', '/..', '/ntd', '\\nsd', '/n.', '', '', '/..d', '...u', '/.s', '\\ns'],
      ['/n.d', '', '/.su', '/ns', '/nsd', '', '/ns', '', '/n.u', '/n.u', '\\..u', '\\nsd', '\\..d', '\\n.', '/n.', '', '', '', '', ''],
      ['/n.d', '/w.u', '', '', '', '', '/n.u', '/.s', '/ns', '\\n.', '\\nsd', '/wtd', '\\n.u', '/.s', '\\.s', '', '', '', '/nt', ''],
      ['/n.d', '\\nsd', '', '', '', '', '/n.', '', '', '\\.s', '\\.s', '/nsu', '\\.td', '\\n.d', '\\.s', '/..', '/n.u', '', '\\ns', ''],
      ['/n.d', '\\ns', '/nsd', '', '/..', '', '\\ntd', '/wtu', '', '', '\\n.', '', '', '/nt', '\\nsd', '\\.tu', '/n.', '/n.d', '\\n.d', '/nsu'],
      ['/.su', '', '/nsu', '', '', '\\n.', '\\.s', '', '\\nsu', '\\nsu', '/wtu', '\\.s', '\\w.', '\\ns', '/..u', '', '', '\\.s', '/n.u', '/nsd'],
      ['/.sd', '', '', '/nsd', '\\ns', '\\ns', '\\..', '', '/ns', '/..d', '/wt', '/ns', '', '', '/.tu', '', '\\.sd', '', '\\n.u', ''],
      ['\\.sd', '/wt', '/..d', '/..u', '/w.', '/n.', '\\n.', '', '', '...d', '\\.sd', '', '\\..u', '\\..d', '', '/n.', '\\n.d', '', '/n.', '\\n.u'],
      ['/..d', '/.s', '/n.d', '/w.d', '', '\\ns', '/.tu', '/n.u', '\\.s', '/n.u', '', '\\n.d', '/nsd', '/.sd', '/nsu', '', '/..', '/n.', '\\.s', '\\..u'],
      ['/n.', '/.su', '', '', '\\ns', '\\nsd', '', '', '/wtd', '/n..', '\\ns', '/.su', '', '\\.s', '/.sd', '', '/n.', '/..u', '', '/n.'],
      ['', '/n.d', '/n.', '\\.su', '/..', '', '/.s', '', '\\nsd', '/ns.', '\\n.u', '/nsu', '/..', '\\..', '/w.u', '/..d', '\\..u', '\\nsd', '\\n.u', '/n.u'],
    ]
    return comments,


if __name__ == "__main__":
    app.run()
