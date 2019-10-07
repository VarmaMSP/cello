import classNames from 'classnames'
import React from 'react'
import { connect } from 'react-redux'
import { getViewportSize } from 'selectors/browser/viewport'
import { AppState } from 'store'
import { ViewportSize } from 'types/app'

interface StateToProps {
  viewPortSize: ViewportSize
}

interface OwnProps {
  // No of Rows
  //  - If Specified The grid is truncated to specfied no of rows
  //  - Defaults to undefined
  rows?: { [key in ViewportSize]: number } | number

  // No of Columns
  //  - If a object is provided grid will be responsive otherwise rows
  //    will overflow in x-axis with each childs width set to provided width
  //  - Defaults to { LG: 7 , MD: 6, SM: 3 }
  cols?: { [key in ViewportSize]: number } | number

  // Optional css class assigned to each row
  className?: string

  // Optional css class assigned to each child wrapper
  classNameChild?: string

  // percentage of container width that is spread amoung children in each row
  //  - default to 12%
  totalRowSpacing?: { [key in ViewportSize]: number } | number

  // Children
  children: JSX.Element[]
}

const Grid: React.SFC<StateToProps & OwnProps> = ({
  rows,
  cols = { LG: 7, MD: 6, SM: 3 },
  className = '',
  classNameChild = '',
  totalRowSpacing = 12,
  children,
  viewPortSize,
}) => {
  let itemsOverflow = typeof cols === 'number' && !!cols
  let itemsPerRow = typeof cols === 'object' ? cols[viewPortSize] : cols
  let itemWidth =
    typeof totalRowSpacing === 'object'
      ? (100 - totalRowSpacing[viewPortSize]) / itemsPerRow
      : (100 - totalRowSpacing) / itemsPerRow

  const items = children.map((child, i) => (
    <div
      key={i}
      className={classNames('flex-none', classNameChild)}
      style={{ width: !itemsOverflow ? `${itemWidth}%` : undefined }}
    >
      {child}
    </div>
  ))

  // Add Placeholders if required
  let len = items.length
  if (len % itemsPerRow > 0) {
    let placeholderCount = itemsPerRow - (len % itemsPerRow)
    while (placeholderCount--) {
      items.push(
        <div
          key={len + placeholderCount}
          className="flex-none"
          style={{ width: !itemsOverflow ? `${itemWidth}%` : undefined }}
        />,
      )
    }
  }

  let rowsJsx: JSX.Element[] = []
  for (let i = 0; i < items.length; i += itemsPerRow) {
    rowsJsx.push(
      <div
        key={i + itemsPerRow}
        className={classNames('flex justify-around', className, {
          'flex-wrap-none': itemsOverflow,
          'overflow-x-auto': itemsOverflow,
        })}
      >
        {items.slice(i, i + itemsPerRow)}
      </div>,
    )
  }

  if (typeof rows === 'object') {
    return <div>{rows ? rowsJsx.slice(0, rows[viewPortSize]) : rowsJsx}</div>
  }
  if (typeof rows === 'number') {
    return <div>{rows ? rowsJsx.slice(0, rows) : rowsJsx}</div>
  }
  return <div>{rowsJsx}</div>
}

function mapStateToProps(state: AppState): StateToProps {
  return { viewPortSize: getViewportSize(state) }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  Grid,
)
