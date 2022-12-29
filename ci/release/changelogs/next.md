#### Features 🚀

- Tooltips can be set on shapes. See [https://d2lang.com/tour/tooltips](https://d2lang.com/tour/interactive). [#548](https://github.com/terrastruct/d2/pull/548)
- Links can be set on shapes. See [https://d2lang.com/tour/tooltips](https://d2lang.com/tour/interactive). [#548](https://github.com/terrastruct/d2/pull/548)

#### Improvements 🧹

- Now the `width` and `height` attributes are no longer restricted to images and can be applied to non-container shapes. [#498](https://github.com/terrastruct/d2/pull/498)

#### Bugfixes ⛑️

- Restricts where `near` key constant values can be used, with good error messages, instead of erroring (e.g. setting `near: top-center` on a container would cause bad layouts or error). [#538](https://github.com/terrastruct/d2/pull/538)
- Fixed rendering classes and tables with empty headers. [#498](https://github.com/terrastruct/d2/pull/498)
