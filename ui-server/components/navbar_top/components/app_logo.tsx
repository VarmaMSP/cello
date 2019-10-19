import { iconMap } from 'components/icon'
import React from 'react'

const AppLogo: React.SFC<{}> = () => {
  const LogoIcon = iconMap['phenopod']

  return <LogoIcon className="mx-auto" />
}

export default AppLogo
