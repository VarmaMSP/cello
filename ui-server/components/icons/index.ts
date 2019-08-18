import React from 'react'
import SvgArrowLeft from './ArrowLeft';
import SvgCheveronUp from './CheveronUp';
import SvgExplore from './Explore';
import SvgFeed from './Feed';
import SvgFastForward from './FastForward';
import SvgFastRewind from './FastRewind';
import SvgHeart from './Heart';
import SvgPlay from './Play';
import SvgSearch from './search';
import SvgUserSolidCircle from './UserSolidCircle';

export type Icon
  = 'arrow-left'
  | 'cheveron-up'
  | 'explore'
  | 'fast-forward'
  | 'fast-rewind'
  | 'feed'
  | 'heart'
  | 'play'
  | 'search'
  | 'user-solid-circle'

export const iconMap: {[key in Icon]: React.SFC<{className: string}>} = {
  'arrow-left': SvgArrowLeft,
  'cheveron-up': SvgCheveronUp,
  'explore': SvgExplore,
  'fast-forward': SvgFastForward,
  'fast-rewind': SvgFastRewind,
  'feed': SvgFeed,
  'heart': SvgHeart,
  'play': SvgPlay,
  'search': SvgSearch,
  'user-solid-circle': SvgUserSolidCircle,
}
