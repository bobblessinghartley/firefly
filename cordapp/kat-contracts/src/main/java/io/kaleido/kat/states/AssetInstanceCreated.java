package io.kaleido.kat.states;

import io.kaleido.kat.contracts.KatContract;
import net.corda.core.contracts.BelongsToContract;
import net.corda.core.contracts.ContractState;
import net.corda.core.identity.AbstractParty;
import net.corda.core.identity.Party;
import org.jetbrains.annotations.NotNull;

import java.util.List;

@BelongsToContract(KatContract.class)
public class AssetInstanceCreated implements AssetEventState {
    private final String assetInstanceID;
    private final String assetDefinitionID;
    private final Party author;
    private final String contentHash;

    public AssetInstanceCreated(String assetInstanceID, String assetDefinitionID, Party author, String contentHash) {
        this.assetInstanceID = assetInstanceID;
        this.assetDefinitionID = assetDefinitionID;
        this.author = author;
        this.contentHash = contentHash;
    }

    @NotNull
    @Override
    public List<AbstractParty> getParticipants() {
        return List.of(author);
    }

    @Override
    public String toString() {
        return String.format("AssetInstanceCreated(assetInstanceID=%s, assetDefinitionID=%s, author=%s, contentHash=%s)", assetInstanceID, assetDefinitionID, author, contentHash);
    }

    @Override
    public Party getAuthor() {
        return author;
    }
}
