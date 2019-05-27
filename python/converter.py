#!/usr/bin/env python3
# -*- coding:utf-8 -*-
u"""
@author ygidtu@gmail.com
@since 2019.04.29

Convert markdown to html
Convert html to pdf
"""
import os
import re
import sys
from argparse import ArgumentParser

import pdfkit

from jinja2 import Environment, PackageLoader, select_autoescape
from markdown2 import markdown

__dir__ = os.path.dirname(os.path.abspath(__file__))


class Md2Html(object):

    env = Environment(
        loader=PackageLoader('converter', 'templates'),
        autoescape=select_autoescape(['html', 'xml'])
    )

    def __init__(self, input_file, output_file, css_file):
        u"""
        Convert markdown to html using jinja2 and markdown2
        :param input_file:
        :param output_file:
        :param css_file:
        """
        assert os.path.exists(input_file)

        self.input_file = os.path.abspath(input_file)
        self.output_file = os.path.abspath(output_file)

        if self.output_file:
            if not os.path.exists(os.path.dirname(self.output_file)):
                os.makedirs(os.path.dirname(self.output_file))

            if not css_file or not os.path.exists(css_file):
                raise FileNotFoundError("Css {0} not found".format(css_file))

        self.css_file = os.path.abspath(css_file)

    def convert(self):
        u"""
        Convert
        :return: converted html string
        """
        with open(self.css_file) as r:
            css = r.read()

        with open(self.input_file) as r:
            body = r.read()

        body = markdown(body, extras=["tables"])

        template = self.env.get_template('typora.html')

        res = template.render(
            title=os.path.basename(self.input_file),
            body=body,
            css=css
        )

        if self.output_file:
            with open(self.output_file, "w+") as w:
                w.write(res)

        return res


def main(args):
    u"""
    main entry
    :param args:
    :return:
    """
    print(args)

    html = args.html
    if args.md:
        if not args.html:
            html = re.sub(r"(markdown|md)$", "html", args.md, re.I)

        Md2Html(args.md, html, args.css).convert()

    if args.pdf:
        if not html:
            html = args.html

        assert(os.path.exists(html)), "html file does not exists"

        pdfkit.from_file(html, output_path=args.pdf)

        if not args.html and os.path.exists(html):
            os.remove(html)


if __name__ == '__main__':
    parser = ArgumentParser(description="Convert markdown to html and convert html to pdf")
    parser.add_argument(
        "-md",
        help="Path to markdown file",
        type=str
    )

    parser.add_argument(
        "-html",
        help="Path to html file",
        type=str
    )

    parser.add_argument(
        "-pdf",
        help="Path to pdf file",
        type=str
    )

    parser.add_argument(
        "--css",
        help="Path to customized css file",
        default=os.path.join(__dir__, "css/basic.css"),
        type=str
    )

    if len(sys.argv) <= 1:
        parser.print_help()
        exit(0)

    try:
        args = parser.parse_args(sys.argv[1:])
        main(args)
    except ArgumentParser as err:
        print(err)
