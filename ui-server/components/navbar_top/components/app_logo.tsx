import { iconMap } from 'components/icon'
import React from 'react'

const AppLogo: React.SFC<{}> = () => {
  const LogoIcon = iconMap['phenopod']

  return <LogoIcon className="w-14 h-14 mx-auto" />
}

export default AppLogo
