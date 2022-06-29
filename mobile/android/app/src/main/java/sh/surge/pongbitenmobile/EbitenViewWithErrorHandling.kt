package sh.surge.pongbitenmobile

import sh.surge.pongbiten.mobile.EbitenView

import android.content.Context
import android.util.AttributeSet

internal class EbitenViewWithErrorHandling : EbitenView {
    constructor(context: Context?) : super(context)
    constructor(context: Context?, attributeSet: AttributeSet?) : super(context, attributeSet)

    override fun onErrorOnGameUpdate(e: Exception) {
        // You can define your own error handling e.g., using Crashlytics.
        // e.g., Crashlytics.logException(e);
        super.onErrorOnGameUpdate(e)
    }
}