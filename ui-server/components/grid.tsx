import React, { cloneElement } from 'react'
import { connect } from 'react-redux'
import { getViewportSize } from 'selectors/window'
import { AppState } from 'store'
import { ViewportSize } from 'types/app'

interface StateToProps {
  viewPortSize: ViewportSize
}

interface OwnProps {
  cols: { [key in ViewportSize]: number }
  rows?: { [key in ViewportSize]: number }
  children: JSX.Element[]
}

const Grid: React.FC<StateToProps & OwnProps> = ({
  rows,
  cols,
  children,
  viewPortSize,
}) => {
  let itemWidth = 100 / cols[viewPortSize]
  let itemsPerRow = cols[viewPortSize]

  const items = children.map((child) =>
    cloneElement(child, {
      style: { width: `${itemWidth}%` },
    }),
  )

  // Add Placeholders if required
  let len = items.length
  if (len % itemsPerRow > 0) {
    let placeholderCount = itemsPerRow - (len % itemsPerRow)
    while (placeholderCount--) {
      items.push(
        <div
          key={placeholderCount.toString()}
          className="flex-none"
          style={{ width: `${itemWidth}%` }}
        />,
      )
    }
  }

  // Rows
  let rowsJsx: JSX.Element[] = []
  for (let i = 0; i < items.length; i += itemsPerRow) {
    rowsJsx.push(
      <div key={i.toString()} className="flex">
        {items.slice(i, i + itemsPerRow)}
      </div>,
    )
  }

  return <div>{!!rows ? rowsJsx.slice(0, rows[viewPortSize]) : rowsJsx}</div>
}

function mapStateToProps(state: AppState): StateToProps {
  return { viewPortSize: getViewportSize(state) }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  Grid,
)
