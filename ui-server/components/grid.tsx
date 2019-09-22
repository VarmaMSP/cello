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
  rows?: number

  // No of Columns
  //  - If a object is provided grid will be responsive otherwise rows
  //    will overflow in x-axis with each childs width set to provided width
  //  - Defaults to { LG: 7 , MD: 6, SM: 3 }
  cols?: { [key in ViewportSize]: number } | number

  // Optional css class assigned to each row
  className?: string

  // Optional css class assigned to each child wrapper
  classNameChild?: string

  // Children
  children: JSX.Element[]
}

const Grid: React.SFC<StateToProps & OwnProps> = ({
  rows,
  cols = { LG: 7, MD: 6, SM: 3 },
  className = '',
  classNameChild = '',
  children,
  viewPortSize,
}) => {
  let overflowRow = typeof cols === 'number' && !!cols
  let itemsPerRow = typeof cols === 'object' ? cols[viewPortSize] : cols

  // Add Placeholders if required
  if (children.length % itemsPerRow > 0) {
    let placeholderCount = itemsPerRow - (children.length % itemsPerRow)
    while (placeholderCount--) {
      children.push(<div />)
    }
  }

  let rowsJsx: JSX.Element[] = []
  for (let i = 0; i < children.length; i += itemsPerRow) {
    rowsJsx.push(
      <div
        key={i + itemsPerRow}
        className={classNames('flex justify-between', className, {
          'flex-wrap-none': overflowRow,
          'overflow-x-auto': overflowRow,
        })}
      >
        {children.slice(i, i + itemsPerRow).map((item, j) => (
          <div
            key={j}
            className={classNames('flex-none', classNameChild)}
            style={{ width: !overflowRow ? `${88 / itemsPerRow}%` : undefined }}
          >
            {item}
          </div>
        ))}
      </div>,
    )
  }

  return <div>{rows ? rowsJsx.slice(0, rows) : rowsJsx}</div>
}

function mapStateToProps(state: AppState): StateToProps {
  return { viewPortSize: getViewportSize(state) }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  Grid,
)
