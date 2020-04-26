var unified = require('unified')
var markdown = require('remark-parse')
var remark2rehype = require('remark-rehype')
var rehype2remark = require('rehype-remark')
// var doc = require("rehype-document");
var format = require('rehype-format')
var html = require('rehype-stringify')
// var report = require("vfile-reporter");
// var vfile = require('to-vfile')
// var report = require('vfile-reporter')
var htmlparse = require('rehype-parse')

var stringify = require('remark-stringify')
// var frontmatter = require('remark-frontmatter')

export function md2html(md) {
  let processor = unified()
    .use(markdown)
    .use(remark2rehype)
    .use(format)
    .use(html)
  let htmltext = processor.processSync(md).toString()
  return htmltext
}

export function html2md(htmltext) {
  let processor = unified()
    .use(htmlparse)
    .use(rehype2remark)
    .use(stringify)
  let md = processor.processSync(htmltext).toString()
  return md
}

export function html2txt(htmltext) {
  let processor = unified()
    .use(htmlparse)
    .use(rehype2remark)
    .use(stringify)
  let md = processor.processSync(htmltext).toString()
  return md
}
