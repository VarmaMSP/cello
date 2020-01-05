import classNames from 'classnames'
import Router from 'next/router'
import React from 'react'
import { SearchResultType, SearchSortBy } from 'types/search'

export interface StateToProps {
  searchBarText: string
}

export interface OwnProps {
  resultType: SearchResultType
  sortBy: SearchSortBy
}

const SearchResultsFilter: React.FC<StateToProps & OwnProps> = (props) => {
  const onResultTypeChange = (t: SearchResultType) => {
    if (props.resultType !== t) {
      Router.push(
        {
          pathname: '/results',
          query: {
            query: props.searchBarText,
            resultType: t,
            sortBy: props.sortBy,
          },
        },
        `/results?query=${props.searchBarText}&type=${t}&sort_by=${props.sortBy}`,
      )
    }
  }

  const onSortByChange = (s: SearchSortBy) => {
    if (props.sortBy !== s) {
      Router.push(
        {
          pathname: '/results',
          query: {
            query: props.searchBarText,
            resultType: props.resultType,
            sortBy: s,
          },
        },
        `/results?query=${props.searchBarText}&type=${props.resultType}&sort_by=${s}`,
      )
    }
  }

  return (
    <div className="flex items-center mb-8 justify-between">
      <div className="flex flex-initial w-3/5 border-b">
        {(['episode', 'podcast'] as SearchResultType[]).map((t) => (
          <div className="w-20 mr-2 text-center">
            <div
              className={classNames(
                'block px-3 py-1 text-sm capitalize leading-loose tracking-wider',
                {
                  'cursor-default': props.resultType === t,
                  'cursor-pointer': props.resultType !== t,
                },
              )}
              onClick={() => onResultTypeChange(t)}
            >
              {`${t}s`}
            </div>
            <div
              className={classNames('h-1 w-20 rounded-full', {
                'bg-green-600': props.resultType === t,
              })}
            />
          </div>
        ))}
      </div>
      <div>
        <select
          className="form-select text-sm tracking-wider w-36"
          onChange={(e) => onSortByChange(e.target.value as SearchSortBy)}
        >
          <option value="relevance">{'Relevance'}</option>
          {props.resultType === 'episode' && (
            <option value="publish_date">{'Date Published'}</option>
          )}
        </select>
      </div>
    </div>
  )
}

export default SearchResultsFilter
