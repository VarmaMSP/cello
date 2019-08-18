import React from 'react';
import classnames from 'classnames'
import IconButton from './buttons/icon_button';

interface Props {
  hide: Boolean;
}

const SearchBarMobile:React.SFC<Props> = ({hide}) => {
  if (hide) return <></>

  return <div className="relative flex h-full w-full px-2 py-1" >
    <IconButton className="absolute inset-y-0 left-0 w-5 ml-4" icon="arrow-left"/>
    <IconButton className="absolute inset-y-0 right-0 w-5 mr-4" icon="search"/>
    <input 
      className={classnames(
        "w-full h-full pl-10 pr-4 py-1 bg-gray-200 text-gray-800 placeholder-gray-600 border border-transparent rounded-lg",
        "appearance-none focus:outline-none focus:bg-gray-100 focus:border-blue-400"
      )}
      placeholder="Search podcasts"
    />
  </div>
}

export default SearchBarMobile