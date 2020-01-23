import React from 'react'
import { connect } from 'react-redux'
import { getViewportSize } from 'selectors/window'
import { AppState } from 'store'
import { ViewportSize } from 'types/app'

interface StateToProps {
  viewPortSize: ViewportSize
}

interface OwnProps {
  children: JSX.Element | JSX.Element[]
}

const PageLayout: React.FC<StateToProps & OwnProps> = ({
  viewPortSize,
  children,
}) => {
  if (viewPortSize === 'SM') {
    return Array.isArray(children) ? (
      <div className="flex flex-col px-3 pt-20 pb-40">
        {children[0]}
        {children[1]}
      </div>
    ) : (
      <div className="px-3 pt-20 pb-40">{children}</div>
    )
  }

  return Array.isArray(children) ? (
    <div
      className="flex flex-row pr-6"
      style={{ paddingLeft: '18rem', paddingBottom: '16rem' }}
    >
      <div className="w-2/3 mr-5">{children[0]}</div>
      <div className="w-1/3">{children[1]}</div>
    </div>
  ) : (
    <div
      className="pr-6"
      style={{ paddingLeft: '17.5rem', paddingBottom: '16rem' }}
    >
      {children}
    </div>
  )
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    viewPortSize: getViewportSize(state),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  PageLayout,
)
