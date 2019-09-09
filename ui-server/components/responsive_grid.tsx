import React from 'react'
import { connect } from 'react-redux'
import { getScreenWidth } from 'selectors/ui/screen'
import { AppState } from 'store'
import { ScreenWidth } from 'types/app'

interface StateToProps {
  screenWidth: ScreenWidth
}

interface OwnProps {
  rows: number
  children: JSX.Element[]
}

interface Props extends StateToProps, OwnProps {}

const ResponsiveGrid: React.SFC<Props> = (props) => {
  const { rows, children, screenWidth } = props

  let itemsPerRow = 2
  if (screenWidth === 'LG') {
    itemsPerRow = 7
  } else if (screenWidth === 'MD') {
    itemsPerRow = 6
  } else if (screenWidth === 'SM') {
    itemsPerRow = 3
  }

  let placeholderCount = children.length % itemsPerRow
  while (placeholderCount--) {
    children.push(<div />)
  }

  let rowsJsx: JSX.Element[] = []
  for (let i = 0; i < children.length; i += itemsPerRow) {
    rowsJsx.push(
      <div key={i + itemsPerRow} className="flex justify-between mb-6">
        {children.slice(i, i + itemsPerRow).map((item, j) => (
          <div key={j} style={{ width: `${88 / itemsPerRow}%` }}>
            {item}
          </div>
        ))}
      </div>,
    )
  }

  return <div>{rowsJsx.slice(0, rows)}</div>
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    screenWidth: getScreenWidth(state),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  ResponsiveGrid,
)
