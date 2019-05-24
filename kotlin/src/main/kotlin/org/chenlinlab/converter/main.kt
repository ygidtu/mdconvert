package org.chenlinlab.converter

import java.io.File
import org.commonmark.parser.Parser
import org.commonmark.renderer.html.HtmlRenderer

import org.commonmark.ext.gfm.tables.TablesExtension
import org.commonmark.ext.gfm.strikethrough.StrikethroughExtension
import org.commonmark.ext.autolink.AutolinkExtension
import org.commonmark.ext.heading.anchor.HeadingAnchorExtension
import org.commonmark.ext.ins.InsExtension

import org.thymeleaf.TemplateEngine
import org.thymeleaf.templateresolver.ClassLoaderTemplateResolver
import org.thymeleaf.context.Context

import com.openhtmltopdf.pdfboxout.PdfRendererBuilder;



class Converter(private val markdown: File?, private val html: File?, val pdf: File?, css: File?) {

    var cssContent = Converter::class.java.getResource("/basic.css").readText()

    init {
        if (css != null) {
            this.cssContent = css.readBytes().toString()
        }
    }


    fun MdToHtml() {

        // parser markdown
        val extensions = arrayListOf(
            TablesExtension.create(),
            StrikethroughExtension.create(),
            AutolinkExtension.create(),
            HeadingAnchorExtension.create(),
            InsExtension.create()
        )

        val parser = Parser.builder()
            .extensions(extensions)
            .build()

        val renderer = HtmlRenderer
            .builder()
            .extensions(extensions)
            .build()

        val content = renderer.render(
            parser.parseReader(this.markdown!!.inputStream().bufferedReader())
        )

        // thymeleaf engine
        val engine = TemplateEngine()
        val resolver = ClassLoaderTemplateResolver()

        resolver.suffix = ".html"
        resolver.prefix = ""

        engine.setTemplateResolver(resolver)

        // prepare variables
        val context = Context()

        context.setVariable("title", this.markdown.name)
        context.setVariable("body", content)
        context.setVariable("css", this.cssContent)

        engine.process("typora", context, this.html!!.outputStream().bufferedWriter())
    }

    fun HtmlToPdf() {
        val parser = PdfRendererBuilder()

        parser.useFastMode()

        parser.withFile(this.html!!)

        this.pdf!!.outputStream().use {
            parser.toStream(it)

            parser.run()
        }
    }
}



fun main(args: Array<String>) {
    println("Hello")
    val t = Converter(
        File("/Volumes/WD/CloudStation/lung_cancer_10X/讨论/2019.05.17/20190517.md"),
        File("/Volumes/WD/CloudStation/lung_cancer_10X/讨论/2019.05.17/20190517_java.html"),
        File("/Volumes/WD/CloudStation/lung_cancer_10X/讨论/2019.05.17/20190517_java.pdf"),
        null
    )

    t.MdToHtml()
    t.HtmlToPdf()

}