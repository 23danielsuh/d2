#### Features 🚀

- PDF exports support external links [#891](https://github.com/terrastruct/d2/issues/891)
- PDF exports support internal links [#891](https://github.com/terrastruct/d2/issues/966)
- `border-radius` is now supported on connections (ELK and TALA only, since Dagre uses curves). [#913](https://github.com/terrastruct/d2/pull/913)

#### Improvements 🧹

- SVGs are fit to top left by default to avoid issues with zooming. [#954](https://github.com/terrastruct/d2/pull/954)
- Person shapes now have labels below them and don't need to expand as much. [#960](https://github.com/terrastruct/d2/pull/960)

#### Bugfixes ⛑️

- Fixes a regression where PNG backgrounds could be cut off in the appendix. [#941](https://github.com/terrastruct/d2/pull/941)
- Fixes zooming not working in watch mode. [#944](https://github.com/terrastruct/d2/pull/944)
- [API] Fixes `DeleteIDDeltas` giving duplicate deltas in rare cases. [#957](https://github.com/terrastruct/d2/pull/957)
