<template>
  <div class="icon" :class="[`icon--${name}`, `icon--${size}`, { 'has-align-fix': fixAlign }]">
    <component :is="`Icon__${name}`"></component>
  </div>
</template>

<script>
import Icon__add_col_after from '@/assets/icons/add_col_after.svg'
import Icon__add_col_before from '@/assets/icons/add_col_before.svg'
import Icon__add_row_after from '@/assets/icons/add_row_after.svg'
import Icon__add_row_before from '@/assets/icons/add_row_before.svg'
import Icon__bold from '@/assets/icons/bold.svg'
import Icon__checklist from '@/assets/icons/checklist.svg'
import Icon__code from '@/assets/icons/code.svg'
import Icon__combine_cells from '@/assets/icons/combine_cells.svg'
import Icon__delete_col from '@/assets/icons/delete_col.svg'
import Icon__delete_row from '@/assets/icons/delete_row.svg'
import Icon__delete_table from '@/assets/icons/delete_table.svg'
import Icon__github from '@/assets/icons/github.svg'
import Icon__hr from '@/assets/icons/hr.svg'
import Icon__image from '@/assets/icons/image.svg'
import Icon__italic from '@/assets/icons/italic.svg'
import Icon__link from '@/assets/icons/link.svg'
import Icon__mention from '@/assets/icons/mention.svg'
import Icon__ol from '@/assets/icons/ol.svg'
import Icon__paragraph from '@/assets/icons/paragraph.svg'
import Icon__quote from '@/assets/icons/quote.svg'
import Icon__redo from '@/assets/icons/redo.svg'
import Icon__remove from '@/assets/icons/remove.svg'
import Icon__strike from '@/assets/icons/strike.svg'
import Icon__table from '@/assets/icons/table.svg'
import Icon__ul from '@/assets/icons/ul.svg'
import Icon__underline from '@/assets/icons/underline.svg'
import Icon__undo from '@/assets/icons/undo.svg'

export default {
    name:"Icon",
    components:{
      Icon__add_col_after,
      Icon__add_col_before,
      Icon__add_row_after,
      Icon__add_row_before,
      Icon__bold,
      Icon__checklist,
      Icon__code,
      Icon__combine_cells,
      Icon__delete_col,
      Icon__delete_row,
      Icon__delete_table,
      Icon__github,
      Icon__hr,
      Icon__image,
      Icon__italic,
      Icon__link,
      Icon__mention,
      Icon__ol,
      Icon__paragraph,
      Icon__quote,
      Icon__redo,
      Icon__remove,
      Icon__strike,
      Icon__table,
      Icon__ul,
      Icon__underline,
      Icon__undo,
    },
    props: {
        name: {
          type: String,
          required: true
        },
        size: {
            default: 'normal',
        },
        modifier: {
            default: null,
        },
        fixAlign: {
            default: true,
        },
    },
}
</script>

<style lang="scss" scoped>
.icon {
    position: relative;
    display: inline-block;
    vertical-align: middle;
    width: 0.8rem;
    height: 0.8rem;
    margin: 0 .3rem;
    top: -.05rem;
    fill: currentColor;
    
    // &.has-align-fix {
    // 	top: -.1rem;
    // }
    
    &__svg {
        display: inline-block;
        vertical-align: top;
        width: 100%;
        height: 100%;
    }
    
    &:first-child {
        margin-left: 0;
    }
    
    &:last-child {
        margin-right: 0;
    }
    
}

// svg sprite
body > svg,
.icon use > svg,
symbol {
    path,
    rect,
    circle,
    g {
        fill: currentColor;
        stroke: none;
    }
    
    *[d="M0 0h24v24H0z"] {
        display: none;
    }
}
</style>
