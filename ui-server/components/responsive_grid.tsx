import classNames from 'classnames'
import React from 'react'
import { connect } from 'react-redux'
import { getScreenWidth } from 'selectors/ui/screen'
import { AppState } from 'store'
import { ScreenWidth } from 'types/app'

interface StateToProps {
  screenWidth: ScreenWidth
}

interface OwnProps {
  // No of Rows
  // - If Specified The grid is truncated to specfied no of rows
  // - Defaults to undefined
  rows?: number

  // Spacing added to bottom of each row
  // default to 6
  rowSpacing?: number

  // No of Columns
  //  - If a object is provided grid will be responsive otherwise rows
  //    will overflow in x-axis with each childs width set to provided width
  //  - Defaults to { LG: 7 , MD: 6, SM: 3 }
  cols?: { [key in ScreenWidth]: number } | number

  // classNames assigned to the child wrapper
  className?: string

  // Children
  children: JSX.Element[]
}

interface Props extends StateToProps, OwnProps {}

const ResponsiveGrid: React.SFC<Props> = ({
  rows,
  rowSpacing = 6,
  cols = { LG: 7, MD: 6, SM: 3 },
  className = '',
  children,
  screenWidth,
}) => {
  let overflowRow = typeof cols === 'number' && !!cols
  let itemsPerRow = typeof cols === 'object' ? cols[screenWidth] : cols

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
        className={classNames('flex justify-between', `mb-${rowSpacing}`, {
          'flex-wrap-none': overflowRow,
          'overflow-x-auto': overflowRow,
        })}
      >
        {children.slice(i, i + itemsPerRow).map((item, j) => (
          <div
            key={j}
            className={classNames('flex-none', className)}
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
  return { screenWidth: getScreenWidth(state) }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  ResponsiveGrid,
)
