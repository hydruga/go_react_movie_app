import React, { Component } from 'react';

export default class Categories extends Component {
    render(){
        return <h2>Category: {this.props.title}</h2>
        // Receiving props from Route called "title"
    }
}