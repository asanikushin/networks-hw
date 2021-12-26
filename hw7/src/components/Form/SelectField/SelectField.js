import React, { Component } from 'react'

import { map } from 'underscore'

import cn from 'classnames'
import PropTypes from 'prop-types'
import {Label, Input, FormGroup} from 'reactstrap'

import './SelectField.scss'

/**
 * options = [{ value, text }]
 */

export default class SelectField extends Component {

    static propTypes = {
        name: PropTypes.string,
        label: PropTypes.string,
        value: PropTypes.string,
        options: PropTypes.array,
        isMultiple: PropTypes.bool,
        className: PropTypes.string,      
        placeholder: PropTypes.string,
        onChange: PropTypes.func
    }

    static defaultProps = {
        type: 'text',
        value: '',
        isMultiple: false,
        onChange: function () {}
    }

    onChange = e => {
        const value = +e.target.value
        const { name, onChange: cb } = this.props
        cb(name, value)
    }

    render () {
        const {
            type,
            name,
            label,
            value,
            options,
            className,
            placeholder
        } = this.props

        return (
            <FormGroup className={cn('SelectField', className)}>
                {label ? (
                    <Label className='SelectField-Label'>
                      {label}
                    </Label>
                ) : null}
                <Input                  
                    name={name}
                    value={value}
                    type="select"
                    placeholder={placeholder}
                    className='SelectField-Input'
                    onChange={this.onChange}>
                    {map(options, o => <option value={o.value}>{o.text}</option>)}
                </Input>
            </FormGroup>
        )
    }
}